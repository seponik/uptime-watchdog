package checker

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/seponik/uptime-watchdog/internal/config"
)

type URLCheckResult struct {
	Endpoint   config.Endpoint // The endpoint that was checked.
	StatusCode int             // The HTTP status code received from the server.
	Delay      time.Duration   // The time it took to get a response.
	Error      error           // Any error that occurred during the check.
}

type Checker struct {
	endpoint config.Endpoint // The endpoint to check.
	client   *http.Client    // HTTP client used to send requests.
}

// NewChecker creates a new Checker instance for the given URL endpoint.
func NewChecker(endpoint config.Endpoint) *Checker {
	return &Checker{
		endpoint: endpoint,
		client: &http.Client{
			Timeout: time.Duration(endpoint.Timeout),
		},
	}
}

// CheckURL sends a GET request to the endpoint's URL and returns the result.
// Returns a URLCheckResult, which includes the Endpoint, status code, delay, and an error if something went wrong.
func (c *Checker) Check() URLCheckResult {
	start := time.Now()
	response, err := c.client.Get(c.endpoint.URL)
	delay := time.Since(start)

	if err != nil {
		err = parseError(err, c.client.Timeout)

		return URLCheckResult{
			Endpoint: c.endpoint,
			Delay:    delay,
			Error:    err,
		}
	}
	defer response.Body.Close()

	return URLCheckResult{
		Endpoint:   c.endpoint,
		StatusCode: response.StatusCode,
		Delay:      delay,
	}
}

func parseError(err error, timeout time.Duration) error {
	errStr := err.Error()

	if strings.Contains(errStr, "deadline exceeded") {
		return fmt.Errorf("request timed out after %v", timeout)
	}

	if strings.Contains(errStr, "no such host") {
		return fmt.Errorf("host not found")
	}

	if strings.Contains(errStr, "connection refused") {
		return fmt.Errorf("connection refused by server")
	}

	return err
}
