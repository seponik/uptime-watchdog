package util

import (
	"fmt"

	"github.com/seponik/uptime-watchdog/internal/config"
)

func ValidateConfig(cfg *config.Config) error {
	if cfg.WebhookURL == "" {
		return fmt.Errorf("webhook URL is not set in the configuration")
	}

	if len(cfg.Endpoints) == 0 {
		return fmt.Errorf("no endpoints configured to monitor")
	}

	for index, endpoint := range cfg.Endpoints {
		if endpoint.Name == "" {
			return fmt.Errorf("endpoint name is not set for URL: %d", index)
		}

		if endpoint.URL == "" {
			return fmt.Errorf("endpoint URL is not set for endpoint: %d", index)
		}
	}

	return nil
}
