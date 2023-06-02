package config

import (
	"fmt"
	"os"
)

type (
	// Config -.
	Config struct {
		HTTP
		PG
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" env:"HTTP_PORT"`
	}
	// PG -.
	PG struct {
		URL string `env-required:"true"  env:"PG_URL"`
	}
)

func NewConfig() (*Config, error) {
	port := os.Getenv("HTTP_PORT")
	url := os.Getenv("PG_URL")
	fmt.Println(url)
	cfg := &Config{
		HTTP: HTTP{
			Port: port,
		},
		PG: PG{
			URL: url,
		},
	}

	return cfg, nil
}
