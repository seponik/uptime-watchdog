package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Endpoint represents a single URL endpoint to monitor.
type Endpoint struct {
	Name           string   `yaml:"name"`            // Name of the endpoint.
	URL            string   `yaml:"url"`             // URL to monitor.
	Timeout        Duration `yaml:"timeout"`         // Timeout for the request.
	Interval       Duration `yaml:"interval"`        // Interval to check the endpoint.
	ExpectedStatus int      `yaml:"expected_status"` // Expected HTTP status code.
}

// Config holds the configuration for the uptime watchdog.
type Config struct {
	WebhookURL string     `yaml:"webhook_url"` // Slack webhook url for notifications.
	Endpoints  []Endpoint `yaml:"endpoints"`   // List of target endpoints to monitor.
}

type Duration time.Duration

// UnmarshalYAML implements the yaml.Unmarshaler interface for Duration parsing.
func (d *Duration) UnmarshalYAML(value *yaml.Node) error {
	var durationStr string

	if err := value.Decode(&durationStr); err != nil {
		return err
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return err
	}

	*d = Duration(duration)
	return nil
}

// Load reads and parses the config file and returns Config.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
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
