package ipc

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"os"
	"sync"
	"sync/atomic"

	"golang.org/x/sys/unix"
	"google.golang.org/protobuf/proto"
)

const (
	maxFrames   int = 100
	maxFrameLen int = 1024
	maxFds      int = 50
	headerLen   int = 100
)

type Conn struct {
	c      *net.UnixConn
	closed atomic.Bool
	sendMu sync.Mutex
	readMu sync.Mutex
}

func NewConnFromFd(fd int) (*Conn, error) {
	f := os.NewFile(uintptr(fd), "")
	c, err := net.FileConn(f)
	if err != nil {
		return nil, err
	}
	return NewConn(c.(*net.UnixConn))
}

// This function expects SOCK_DGRAM sockets as we rely on it for the
// message boundaries.
func NewConn(c *net.UnixConn) (*Conn, error) {
	if c == nil || c.LocalAddr() == nil {
		return nil, fmt.Errorf("ipc: nil unix conn")
	}

	if network := c.LocalAddr().Network(); network != "unixgram" {
		return nil, fmt.Errorf(
			"ipc: works only with unixgram connections, but got %v", network)
	}

	return &Conn{c: c}, nil
}

// Returns the number of payload bytes sent, not including the header or oob data
func (conn *Conn) sendFrame(frame int, payload []byte, more bool, fds ...int) (n int, err error) {
	var oob []byte
	if len(fds) > 0 {
		oob = unix.UnixRights(fds...)
	}

	h := &Header{
		More:       more,
		FrameIndex: uint32(frame),
		Payload:    payload,
	}

	p, err := proto.Marshal(h)
	if err != nil {
		return 0, err
	}

	if plen := len(p); plen > maxFrameLen {
		return 0, fmt.Errorf(
			"ipc: attempted to send message (%d bytes) exceeding max frame size (%d bytes)",
			plen, maxFrameLen)
	}

	n, _, err = conn.c.WriteMsgUnix(p, oob, nil)
	if err != nil {
		return 0, err
	}

	if expected := len(p); n != expected {
		return 0, fmt.Errorf(
			"ipc: partial message sent (%d bytes), expected %d bytes",
			n, expected)
	}

	return len(payload), nil
}

func (conn *Conn) Send(m *IpcMessage) error {
	return conn.SendWithFds(m, []int{}...)
}

func (conn *Conn) SendWithFds(m *IpcMessage, fds ...int) error {
	if conn.closed.Load() {
		return fmt.Errorf("ipc: this conn is closed")
	}

	conn.sendMu.Lock()
	defer conn.sendMu.Unlock()

	if nFds := len(fds); nFds > maxFds {
		return fmt.Errorf("ipc: fds exceed limit (%v), got %v", maxFds, nFds)
	}

	payload, err := proto.Marshal(m)
	pLeft := len(payload)
	if err != nil {
		return err
	}

	if pLeft > ((maxFrameLen - headerLen) * maxFrames) {
		return fmt.Errorf(
			"ipc: message will exceed max frames limit (%v)",
			maxFrames)
	}

	fdsToSend := fds
	var pSent int
	var nFrames int
	for nFrames == 0 || pLeft > 0 {
		// Leave room for the header
		chunkLen := min((maxFrameLen - headerLen), pLeft)
		chunk := payload[pSent : pSent+chunkLen]

		n, err := conn.sendFrame(nFrames, chunk, (chunkLen < pLeft), fdsToSend...)
		if err != nil {
			// Just close this connection, or the stream will get
			// out of sync and we currently have no recovery
			e := conn.Close()
			if e != nil {
				return errors.Join([]error{e, err}...)
			}
			return err
		}

		pSent += n
		pLeft -= n

		// Send fds with the first frame only
		fdsToSend = nil
		nFrames++
	}

	return nil
}

// The caller owns any returned file descriptors. All received descriptors are closed
// on error.
func (conn *Conn) Recv(m *IpcMessage) (fds []int, err error) {
	conn.readMu.Lock()
	defer conn.readMu.Unlock()

	if conn.closed.Load() {
		return nil, fmt.Errorf("ipc: this conn is closed")
	}

	var payload bytes.Buffer
	var receivedFds []int
	var nFrames uint32
	var more bool = true
	chunk := make([]byte, maxFrameLen)
	oob := make([]byte, unix.CmsgSpace(maxFds*4))

	for more {
		n, oobn, flags, _, err := conn.c.ReadMsgUnix(chunk, oob)
		if err != nil {
			return nil, conn.errAndClose(err, receivedFds)
		}

		if flags&unix.MSG_CTRUNC != 0 {
			return nil, conn.errAndClose(
				fmt.Errorf("ipc: received truncated control message"),
				receivedFds)
		}

		if flags&unix.MSG_TRUNC != 0 {
			return nil, conn.errAndClose(
				fmt.Errorf("ipc: received truncated payload message"),
				receivedFds)
		}

		// We only send fds in the first frame
		if oobn > 0 {
			scms, err := unix.ParseSocketControlMessage(oob[:oobn])
			if err != nil {
				return nil, conn.errAndClose(err, receivedFds)
			}

			for _, scm := range scms {
				if scm.Header.Level != unix.SOL_SOCKET || scm.Header.Type != unix.SCM_RIGHTS {
					// Skip this message as it has no descriptors
					continue
				}

				pfds, err := unix.ParseUnixRights(&scm)
				if err != nil {
					return nil, conn.errAndClose(err, receivedFds)
				}

				for _, fd := range pfds {
					unix.CloseOnExec(fd)
				}

				receivedFds = append(receivedFds, pfds...)
			}

			if len(receivedFds) > maxFds || nFrames > 0 {
				return nil, conn.errAndClose(
					fmt.Errorf(
						"ipc: exceeded maxFds (%d); we only allow on first frame",
						maxFds),
					receivedFds)
			}
		}

		h := &Header{}
		err = proto.Unmarshal(chunk[:n], h)
		if err != nil {
			return nil, conn.errAndClose(err, receivedFds)
		}

		// Check for unexpected/out-of-order frames
		if h.FrameIndex != nFrames {
			return nil, conn.errAndClose(
				fmt.Errorf("ipc: unexpected frame index, expected %v, got %v",
					nFrames, h.FrameIndex),
				receivedFds)
		}

		more = h.More
		_, err = payload.Write(h.Payload)
		if err != nil {
			return nil, conn.errAndClose(
				fmt.Errorf("ipc: could not append payload chunk: %w", err),
				receivedFds)
		}

		nFrames++
		if more && nFrames >= uint32(maxFrames) {
			return nil, conn.errAndClose(
				fmt.Errorf("ipc: exceeded max frames (%d) on recv", maxFrames),
				receivedFds)
		}
	}

	// We should have a full payload and fds (if there were any)
	err = proto.Unmarshal(payload.Bytes(), m)
	if err != nil {
		return nil, conn.errAndClose(
			fmt.Errorf("ipc: error decoding payload: %w", err),
			receivedFds)
	}

	return receivedFds, nil
}

func (conn *Conn) errAndClose(e error, fds []int) error {
	var errs []error
	for _, fd := range fds {
		err := unix.Close(fd)
		if err != nil {
			errs = append(errs, err)
		}
	}

	// Also close the connection as the stream is messed up
	// and we currently can't recover from that yet
	cerr := conn.Close()
	if cerr != nil {
		errs = append(errs, cerr)
	}

	if nerrors := len(errs); nerrors > 0 {
		errs = append(errs, e)
		return errors.Join(errs...)
	}

	return e
}

func (conn *Conn) Close() error {
	if conn.closed.CompareAndSwap(false, true) {
		err := conn.c.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func (conn *Conn) Closed() bool {
	return conn.closed.Load()
}
