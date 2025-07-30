package checker

import (
	"net/http"
	"sync"
	"time"
)

type URLCheckResult struct {
	URL        string
	StatusCode int
	Delay      time.Duration
	Err        error
}

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
			Err:   err,
		}
	}
	defer response.Body.Close()

	return URLCheckResult{
		URL:        url,
		StatusCode: response.StatusCode,
		Delay:      delay,
	}

}

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
