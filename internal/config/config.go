package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config contains settings loaded from config.yaml:
type Config struct {
	Interval   int      `yaml:"interval"`    // How often to run tasks (seconds)
	WebhookURL string   `yaml:"webhook_url"` // Slack webhook url for notifications
	URLs       []string `yaml:"urls"`        // List of target URLs
}

// Load reads and parses the config file and returns Config.
func Load() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
