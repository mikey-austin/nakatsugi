package health

import (
	"context"
	"fmt"
	"net/http"
)

type Checker interface {
	CheckHealth() *CheckResult
}

type CheckVerifier interface {
	Verify(*http.Response) *CheckResult
}

type CheckRequestBuilder interface {
	NewRequest(context.Context) (*http.Request, error)
}

type CheckResultStatus int

const (
	CheckTimedOut CheckResultStatus = iota
	CheckSucceeded
	CheckFailed
	CheckUnknown
)

type CheckResult struct {
	status CheckResultStatus
	error  error
}

func NewTimeoutCheckResult(e error) *CheckResult {
	return &CheckResult{
		status: CheckTimedOut,
		error:  e,
	}
}

func NewErrorCheckResult(e error) *CheckResult {
	return &CheckResult{
		status: CheckFailed,
		error:  e,
	}
}

func NewUnknownCheckResult(e error) *CheckResult {
	return &CheckResult{
		status: CheckUnknown,
		error:  e,
	}
}

func (c *CheckResult) String() string {
	return fmt.Sprintf("<CheckResult: healthy=%v, timedOut=%v", c.IsHealthy(), c.IsTimedOut())
}

func (c *CheckResult) IsHealthy() bool {
	return c.status == CheckSucceeded
}

func (c *CheckResult) IsTimedOut() bool {
	return c.status == CheckTimedOut
}

func (c *CheckResult) IsUnknown() bool {
	return c.status == CheckUnknown
}

func (c *CheckResult) Error() error {
	return c.error
}
