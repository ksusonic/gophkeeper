package config

import (
	"flag"

	"github.com/caarlos0/env/v7"
)

const (
	DefaultAddress     = ":3200"
	DefaultDatabaseDsn = "postgresql://localhost:5432/gophkeeper"
)

type Config struct {
	Server    ServerConfig
	DebugMode bool
}

type ServerConfig struct {
	Address     string `env:"RUN_ADDRESS"`
	DatabaseDsn string `env:"DATABASE_DSN"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	flag.StringVar(&cfg.Server.Address, "a", DefaultAddress, "serve address")
	flag.StringVar(&cfg.Server.DatabaseDsn, "d", DefaultDatabaseDsn, "db connect string")
	flag.BoolVar(&cfg.DebugMode, "debug", false, "debug mode")

	flag.Parse()

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
