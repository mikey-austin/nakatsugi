package health

import (
	"encoding/hex"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/mikey-austin/nakatsugi/internal/config"
)

const (
	// Both digests for "test"
	testMd5       = "098f6bcd4621d373cade4e832627b4f6"
	testSha1      = "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"
	invalidDigest = "deadbeef"
)

func TestHttpChecker(t *testing.T) {
	server, client, addr, port := newTestServer()
	defer server.Close()
	cases := []struct {
		d config.Destination
		h bool
		t bool
	}{
		{d: newDestCodeCheck(addr, port, 200, "/health1"), h: true, t: false},
		{d: newDestCodeCheck(addr, port, 200, "/health2"), h: false, t: false},
		{d: newDestDigestCheck(addr, port, testMd5, "/health1"), h: true, t: false},
		{d: newDestDigestCheck(addr, port, testSha1, "/health1"), h: true, t: false},
		{d: newDestDigestCheck(addr, port, testMd5, "/health2"), h: false, t: false},
		{d: newDestDigestCheck(addr, port, testSha1, "/health2"), h: false, t: false},
		{d: newDestCodeCheck(addr, port, 200, "/timeout"), h: false, t: true},
	}

	for i, c := range cases {
		httpDest, _ := NewHttpDestination(cases[i].d)
		timeout, _ := time.ParseDuration("100ms")
		httpChecker := NewHttpChecker(client, timeout, port, httpDest, httpDest)

		res := httpChecker.CheckHealth(addr)
		if c.h != res.IsHealthy() {
			t.Errorf("TestHttpChecker[%d]: Expected %v check result, got %v", i, c.h, res)
		}

		if c.t != res.IsTimedOut() {
			t.Errorf("TestHttpChecker[%d]: Expected timed out as %v, got %v", i, c.t, res.IsTimedOut())
		}
	}
}

func TestInvalidDigest(t *testing.T) {
	server, _, addr, port := newTestServer()
	defer server.Close()
	_, err := NewHttpDestination(newDestDigestCheck(addr, port, invalidDigest, "/"))
	if err == nil {
		t.Errorf("expected error during construction with an invalid expected digest")
	}
}

// Test data below

func newDestCodeCheck(addr net.IP, port uint16, code int, path string) config.Destination {
	table := config.Table{
		Name: "t1",
		Items: []*config.TableItem{
			&config.TableItem{Type: config.IPv4Host, IP: addr}},
	}

	return config.Destination{
		ToTable: &table,
		Port:    port,
		Check: &config.Check{
			Type: config.CheckTypeHttpCode,
			CheckHttpCode: &config.CheckHttpCode{
				CheckHttpBase: config.CheckHttpBase{Path: path, Https: false},
				Code:          code,
			}},
	}
}

func newDestDigestCheck(addr net.IP, port uint16, digest string, path string) config.Destination {
	expected, _ := hex.DecodeString(digest)
	table := config.Table{
		Name: "t1",
		Items: []*config.TableItem{
			&config.TableItem{Type: config.IPv4Host, IP: addr}},
	}

	return config.Destination{
		ToTable: &table,
		Port:    port,
		Check: &config.Check{
			Type: config.CheckTypeHttpDigest,
			CheckHttpDigest: &config.CheckHttpDigest{
				CheckHttpBase: config.CheckHttpBase{Path: path, Https: false},
				Digest:        expected,
			}},
	}
}

func newTestServer() (*httptest.Server, *http.Client, net.IP, uint16) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/health1":
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("test"))
			case "/health2":
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("other stuff"))
			case "/timeout":
				<-r.Context().Done() // block until client disconnects
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("other stuff"))
			}
		}))

	client := server.Client()
	url, _ := url.Parse(server.URL)
	addr := net.ParseIP("127.0.0.1")
	port, _ := strconv.ParseUint(url.Port(), 10, 16)

	return server, client, addr, uint16(port)
}
