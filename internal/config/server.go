package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env/v8"
)

const (
	DefaultAddress     = ":3200"
	DefaultDatabaseDsn = "postgresql://localhost:5432/gophkeeper"
	DefaultSecretKey   = "do-not-use-this-in-production!"
	DefaultTokenTTL    = time.Minute * 5
)

type Config struct {
	Server    ServerConfig
	Auth      AuthConfig
	DebugMode bool `env:"DEBUG"`
}

type ServerConfig struct {
	Address     string `env:"RUN_ADDRESS"`
	DatabaseDsn string `env:"DATABASE_DSN"`
}

type AuthConfig struct {
	SecretKey string        `env:"SECRET_KEY"`
	TokenTTL  time.Duration `env:"TOKEN_TTL"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	flag.StringVar(&cfg.Server.Address, "a", DefaultAddress, "serve address")
	flag.StringVar(&cfg.Server.DatabaseDsn, "d", DefaultDatabaseDsn, "db connect string")
	flag.DurationVar(&cfg.Auth.TokenTTL, "token-ttl", DefaultTokenTTL, "token time to live")
	flag.BoolVar(&cfg.DebugMode, "debug", false, "debug mode")

	flag.Parse()

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	if len(cfg.Auth.SecretKey) == 0 {
		fmt.Println("WARN: Secret key is not set- using default key. For production use 'export SECRET_KEY=...'")
		cfg.Auth.SecretKey = DefaultSecretKey
	}

	// security reasons
	_ = os.Unsetenv("SECRET_KEY")

	return &cfg, nil
}
