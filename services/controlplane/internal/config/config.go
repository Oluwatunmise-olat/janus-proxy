package config

import (
	"fmt"
	"os"
)

type Config struct {
	HTTPAddr   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func Load() (Config, error) {
	cfg := Config{
		HTTPAddr:   envOrDefault("CONTROLPLANE_ADDR", ":8081"),
		DBHost:     envOrDefault("DB_HOST", ""),
		DBPort:     envOrDefault("DB_PORT", ""),
		DBUser:     envOrDefault("DB_USER", ""),
		DBPassword: envOrDefault("DB_PASSWORD", ""),
		DBName:     envOrDefault("DB_NAME", ""),
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
	if c.DBHost == "" {
		return fmt.Errorf("DB_HOST cannot be empty")
	}
	if c.DBPort == "" {
		return fmt.Errorf("DB_PORT cannot be empty")
	}
	if c.DBUser == "" {
		return fmt.Errorf("DB_USER cannot be empty")
	}
	if c.DBName == "" {
		return fmt.Errorf("DB_NAME cannot be empty")
	}
	return nil
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
