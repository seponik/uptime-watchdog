package checker

import (
	"net/http"
	"sync"
	"time"
)

type URLCheckResult struct {
	URL        string        // The URL that was checked.
	StatusCode int           // The HTTP status code received from the server.
	Delay      time.Duration // The time it took to get a response.
	Error      error         // Any error that occurred during the check.
}

// CheckURL sends a GET request to the given URL to check its status.
// The default timeout is 5 seconds.
// Returns a URLCheckResult, which includes the URL, status code, delay, and an error if something went wrong.
func CheckURL(url string) URLCheckResult {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now()
	response, err := client.Get(url)
	delay := time.Since(start)

	if err != nil {
		return URLCheckResult{
			URL:   url,
			Delay: delay,
			Error: err,
		}
	}
	defer response.Body.Close()

	return URLCheckResult{
		URL:        url,
		StatusCode: response.StatusCode,
		Delay:      delay,
	}

}

// CheckAll concurrently checks the provided URLs with CheckURL function.
// Returns a slice of URLCheckResult containing the results for each URL.
func CheckAll(urls []string) []URLCheckResult {
	if len(urls) == 0 {
		return nil
	}

	results := make([]URLCheckResult, len(urls))
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)

		go func(i int, url string) {
			defer wg.Done()

			results[i] = CheckURL(url)
		}(i, url)
	}

	wg.Wait()

	return results
}
