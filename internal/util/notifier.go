package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/seponik/uptime-watchdog/internal/checker"
	"github.com/seponik/uptime-watchdog/internal/config"
)

type Notifier struct {
	webhookURL string // Slack webhook URL for notifications
}

// NewNotifier creates a new Notifier instance with the given webhook URL.
func NewNotifier(webhookURL string) Notifier {
	return Notifier{webhookURL: webhookURL}
}

// ProcessResult processes the URL check result and sends an alert if necessary.
func (n *Notifier) ProcessResult(result checker.URLCheckResult) {
	printResult(result)
	if result.Error != nil || result.StatusCode != result.Endpoint.ExpectedStatus {
		if err := sendAlert(n.webhookURL, result.Endpoint); err != nil {
			log.Printf("[ERR]  Failed to send alert: %v", err)
		}
	}
}

// printResults logs each Endpoints's check result: UP if OK, DOWN if ERROR or Unexpected Status code.
func printResult(result checker.URLCheckResult) {
	if result.Error != nil || result.StatusCode != result.Endpoint.ExpectedStatus {
		log.Printf("[DOWN] %-20s error: %-20v",
			result.Endpoint.Name,
			result.Error,
		)

		return
	}

	log.Printf("[UP]   %-20s status: %-4d (%.2vms)",
		result.Endpoint.Name,
		result.StatusCode,
		result.Delay.Milliseconds(),
	)
}

// sendAlert sends a alert to the given slack webhook if a URL is down.
// Returns an error if something went wrong.
func sendAlert(webhookURL string, endpoint config.Endpoint) error {
	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	alertMessage := fmt.Sprintf("🚨 ALERT: %s is DOWN as of %s UTC.", endpoint.Name, timestamp)

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
