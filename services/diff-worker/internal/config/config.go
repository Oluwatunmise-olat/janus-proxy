package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	TickInterval time.Duration
}

func Load() (Config, error) {
	cfg := Config{
		TickInterval: durationOrDefault("DIFF_WORKER_TICK_INTERVAL", 30*time.Second),
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c Config) Validate() error {
	if c.TickInterval <= 0 {
		return fmt.Errorf("DIFF_WORKER_TICK_INTERVAL must be greater than zero")
	}

	return nil
}

func durationOrDefault(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := time.ParseDuration(value)

	if err != nil {
		return fallback
	}

	return parsed
}
