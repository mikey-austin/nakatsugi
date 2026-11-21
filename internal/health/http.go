package health

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

type HttpChecker struct {
	verifier   CheckVerifier
	reqBuilder CheckRequestBuilder
	client     *http.Client
	timeout    time.Duration
}

func NewHttpChecker(
	client *http.Client,
	checkTimeout time.Duration,
	verifier CheckVerifier,
	reqBuilder CheckRequestBuilder) *HttpChecker {

	return &HttpChecker{
		verifier:   verifier,
		reqBuilder: reqBuilder,
		client:     client,
		timeout:    checkTimeout,
	}
}

func (h *HttpChecker) CheckHealth() *CheckResult {
	// TODO: we may need to be able to cancel in-flight checks
	ctx, _ := context.WithTimeout(context.Background(), h.timeout)
	req, err := h.reqBuilder.NewRequest(ctx)
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
