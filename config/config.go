// Package config provides loading config data from external sources
// like env variables, yaml-files etc.
package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App
		DB
	}

	App struct {
		Port         string        `env:"APP_PORT" env-default:"8080" env-description:"app port (default: 8080)"`
		TokenExpired time.Duration `env:"TOKEN_EXPIRED" env-default:"30m" env-description:"JWT token expired duration (default: 30m)"`
	}

	DB struct {
		Path string `env-required:"true" env:"DB_PATH" env-description:"Path to SQLite db"`
	}
)

// Returns app config loaded from ENV-vars.
func New() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("create config: %w", err)
	}
	return cfg, nil
}
