package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/seponik/uptime-watchdog/internal/checker"
)

func ProcessResults(results []checker.URLCheckResult, webhookURL string) {
	printResults(results)
	for _, result := range results {

		if result.Err != nil || result.StatusCode >= 300 {
			err := sendAlert(webhookURL, result.URL)
			if err != nil {
				log.Printf("Failed to send alert for %s: %v", result.URL, err)
			}
		}

	}
}

func printResults(results []checker.URLCheckResult) {
	for _, result := range results {
		if result.Err != nil {
			log.Printf("[DOWN] %s - error: %v (%.2fs)",
				result.URL,
				result.Err,
				result.Delay.Seconds(),
			)

			continue
		}

		if result.StatusCode >= 300 {
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

func sendAlert(webhookURL, url string) error {
	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	alertMessage := fmt.Sprintf("🚨 ALERT: %s is DOWN as of %s UTC.", url, timestamp)

	payload := map[string]string{"text": alertMessage}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal alert payload: %v", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to send alert: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status code from webhook: %d", resp.StatusCode)
	}

	return nil
}
