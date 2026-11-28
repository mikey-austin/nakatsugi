package ipc

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/sys/unix"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var hbMsg HeartBeatMessage = HeartBeatMessage{Timestamp: timestamppb.Now()}
var sdMsg ShutdownMessage = ShutdownMessage{}

func TestSimpleMessageExchange(t *testing.T) {
	fds, err := unix.Socketpair(unix.AF_LOCAL, unix.SOCK_DGRAM, 0)
	if err != nil {
		t.Fatal(err)
	}

	sideA, err := NewConnFromFd(fds[0])
	if err != nil {
		t.Fatal(err)
	}
	if sideA.Closed() {
		t.Fatalf("socket is closed unexpectedly")
	}

	sideB, err := NewConnFromFd(fds[1])
	if err != nil {
		t.Fatal(err)
	}
	if sideB.Closed() {
		t.Fatalf("socket is closed unexpectedly")
	}

	//
	// Send a message through the socket, in both directions
	//
	in := IpcMessage{Payload: &IpcMessage_HeartBeatMsg{HeartBeatMsg: &hbMsg}}
	err = sideA.Send(&in)
	if err != nil {
		t.Fatal(err)
	}

	// Receive
	var out IpcMessage
	_, err = sideB.Recv(&out)
	if err != nil {
		t.Fatal(err)
	}

	if !proto.Equal(&in, &out) {
		t.Errorf("received unexpected message, got {%v}, expected {%v}", &in, &out)
	}

	//
	// Send a file descriptor in the opposite direction through the socket
	//
	testFilePath := filepath.Join(os.TempDir(), "ipctest.txt")
	os.WriteFile(testFilePath, []byte("hello world"), 0600)
	defer os.Remove(testFilePath)

	testFile, err := os.Open(testFilePath)
	if err != nil {
		t.Fatal(err)
	}
	defer testFile.Close()

	fdToSend := int(testFile.Fd())
	err = sideB.SendWithFds(&in, fdToSend)
	if err != nil {
		t.Fatal(err)
	}

	receivedFds, err := sideA.Recv(&out)
	if err != nil {
		t.Fatal(err)
	}

	if gotFds := len(receivedFds); gotFds != 1 {
		t.Fatalf("expected %d fds, got %d", 1, gotFds)
	}

	dupTestFile := os.NewFile(uintptr(receivedFds[0]), "ipcreceived.txt")
	defer dupTestFile.Close()

	buf, err := io.ReadAll(dupTestFile)
	if err != nil {
		t.Fatal(err)
	}

	if got := string(buf); got != "hello world" {
		t.Errorf("expected '%s', got '%v'", "hello world", got)
	}
}
