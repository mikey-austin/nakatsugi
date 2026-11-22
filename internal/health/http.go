package health

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/mikey-austin/nakatsugi/internal/config"
)

type HttpChecker struct {
	verifier   CheckVerifier
	reqBuilder CheckRequestBuilder
	client     *http.Client
	timeout    time.Duration
	port       uint16
}

func NewHttpChecker(
	client *http.Client,
	timeout time.Duration,
	port uint16,
	verifier CheckVerifier,
	reqBuilder CheckRequestBuilder) *HttpChecker {

	return &HttpChecker{
		verifier:   verifier,
		reqBuilder: reqBuilder,
		client:     client,
		timeout:    timeout,
		port:       port,
	}
}

func (h *HttpChecker) CheckHealth(addr net.IP) *CheckResult {
	// TODO: we may need to be able to cancel in-flight checks
	ctx, _ := context.WithTimeout(context.Background(), h.timeout)
	req, err := h.reqBuilder.NewRequest(ctx, addr, h.port)
	if err != nil {
		return NewUnknownCheckResult(err)
	}

	resp, err := h.client.Do(req)
	if e, ok := err.(*url.Error); ok {
		if e.Timeout() {
			return NewTimeoutCheckResult(e)
		}
		return NewErrorCheckResult(e)
	}

	return h.verifier.Verify(resp)
}

type HttpDestination struct {
	scheme   string
	path     string
	hostname string
	verify   func(*http.Response) (*CheckResult, error)
}

func NewHttpRequestBuilder(destination config.Destination) (*HttpDestination, error) {
	dest := &HttpDestination{}
	var base config.CheckHttpBase

	switch destination.Check.Type {
	case config.CheckTypeHttpCode:
		base = destination.Check.CheckHttpCode.CheckHttpBase
		code := destination.Check.CheckHttpCode.Code
		dest.verify = func(resp *http.Response) (*CheckResult, error) {
			return dest.verifyByCode(code, resp)
		}
	case config.CheckTypeHttpDigest:
		base = destination.Check.CheckHttpDigest.CheckHttpBase
		expected := destination.Check.CheckHttpDigest.Digest

		var hash func() hash.Hash
		if len(expected) == md5.Size {
			hash = md5.New
		} else if len(expected) == sha1.Size {
			hash = sha1.New
		} else {
			return nil, fmt.Errorf(
				"expected digest looks wrong; expecting a %d byte MD5 or %d byte SHA1 digest",
				md5.Size, sha1.Size)
		}

		dest.verify = func(resp *http.Response) (*CheckResult, error) {
			return dest.verifyByDigest(hash, expected, resp)
		}
	default:
		return nil, fmt.Errorf("not an http based check")
	}

	dest.hostname = base.Hostname
	dest.path = base.Path
	dest.scheme = "http"
	if base.Https {
		dest.scheme = "https"
	}

	return dest, nil
}

func (h *HttpDestination) Verify(resp *http.Response) *CheckResult {
	res, err := h.verify(resp)
	if err != nil {
		return NewUnknownCheckResult(err)
	}
	return res
}

func (h *HttpDestination) verifyByDigest(
	newHash func() hash.Hash, expected []byte, resp *http.Response) (*CheckResult, error) {

	digest := newHash()
	_, err := io.Copy(digest, resp.Body)
	if err != nil {
		return nil, err
	}

	computed := digest.Sum(nil)
	if bytes.Equal(computed, expected) {
		return NewSuccessfulCheckResult(), nil
	}

	return NewErrorCheckResult(
		fmt.Errorf("health check digest mismatch: expected %x, but got %x", expected, computed)), nil
}

func (h *HttpDestination) verifyByCode(code int, resp *http.Response) (*CheckResult, error) {
	if resp.StatusCode == code {
		return NewSuccessfulCheckResult(), nil
	}
	return NewErrorCheckResult(
		fmt.Errorf("health check http status code mismatch: expected %d but got %d", code, resp.StatusCode)), nil
}

func (h *HttpDestination) NewRequest(ctx context.Context, addr net.IP, port uint16) (*http.Request, error) {
	url := url.URL{
		Scheme: h.scheme,
		Host:   net.JoinHostPort(addr.String(), strconv.Itoa(int(port))),
		Path:   h.path,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	if h.hostname != "" {
		req.Host = h.hostname
	}

	return req, nil
}
