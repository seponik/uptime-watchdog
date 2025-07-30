package util

import (
	"log"
	"net/http"

	"github.com/seponik/uptime-watchdog/internal/checker"
)

func PrintResults(results []checker.URLCheckResult) {
	for _, result := range results {
		if result.Err != nil {
			log.Printf("[DOWN] %s - error: %v (%.2fs)",
				result.URL,
				result.Err,
				result.Delay.Seconds(),
			)

			continue
		}

		if result.StatusCode != http.StatusOK {
			log.Printf("[WARN] %s - status: %d (%.2fs)",
				result.URL,
				result.StatusCode,
				result.Delay.Seconds(),
			)

			continue
		}

		log.Printf("[UP]   %s - status: %d (%.2fs)",
			result.URL,
			result.StatusCode,
			result.Delay.Seconds(),
		)
	}
}
