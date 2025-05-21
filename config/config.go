package config

import "github.com/caarlos0/env/v10"

// Config is the configuration for the application.
type Config struct {
	Environment string `env:"ENVIRONMENT" envDefault:"local"`
}

// ParseConfig parses the configuration for the application.
func ParseConfig[T any](cfg *T) error {
	return env.Parse(cfg)
}
