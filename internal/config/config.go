package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Interval int      `yaml:"interval"`
	URLs     []string `yaml:"urls"`
}

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
