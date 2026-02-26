package config

import (
	"fmt"
	"os"
)

type Config struct {
	HTTPAddr string
}

func Load() (Config, error) {
	cfg := Config{
		HTTPAddr: envOrDefault("CONTROLPLANE_ADDR", ":8081"),
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c Config) Validate() error {
	if c.HTTPAddr == "" {
		return fmt.Errorf("CONTROLPLANE_ADDR cannot be empty")
	}
	return nil
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
