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
		Port            string        `env:"APP_PORT" env-default:"8080"`
		AuthTokenSecret []byte        `env-required:"true" env:"AUTH_TOKEN_SECRET"`
		AuthTokenTTL    time.Duration `env:"AUTH_TOKEN_TTL" env-default:"30m"`
	}

	DB struct {
		Path string `env-required:"true" env:"DB_PATH"`
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
