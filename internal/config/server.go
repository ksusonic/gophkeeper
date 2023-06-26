package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env/v8"
)

const (
	DefaultAddress        = ":3200"
	DefaultDatabaseDsn    = "postgresql://localhost:5432/gophkeeper"
	DefaultSecretKey      = "do-not-use-this-in-production!"
	DefaultTokenTTL       = time.Minute * 15
	DefaultTLSPath        = "cert/server-tls/"
	DefaultSecretKeysPath = "cert/secret-encryption/"
)

type Config struct {
	Server    ServerConfig
	Auth      AuthConfig
	Secrets   SecretsConfig
	DebugMode bool `env:"DEBUG"`
}

type ServerConfig struct {
	Address      string `env:"RUN_ADDRESS"`
	TLSCertsPath string `env:"TLS_PATH"`
	DatabaseDsn  string `env:"DATABASE_DSN"`
}

type AuthConfig struct {
	SaltKey  string        `env:"SALT_KEY"`
	TokenTTL time.Duration `env:"TOKEN_TTL"`
}

type SecretsConfig struct {
	KeysDir string `env:"SECRET_KEYS_DIR"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	flag.StringVar(&cfg.Server.Address, "address", DefaultAddress, "serve address")
	flag.StringVar(&cfg.Server.DatabaseDsn, "dsn", DefaultDatabaseDsn, "db connect string")
	flag.StringVar(&cfg.Server.TLSCertsPath, "tls-path", DefaultTLSPath, "path to tls certs")
	flag.DurationVar(&cfg.Auth.TokenTTL, "token-ttl", DefaultTokenTTL, "token time to live")
	flag.StringVar(&cfg.Secrets.KeysDir, "secret-keys-path", DefaultSecretKeysPath, "path to private and public keys")
	flag.BoolVar(&cfg.DebugMode, "debug", false, "debug mode")

	flag.Parse()

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	if len(cfg.Auth.SaltKey) == 0 {
		fmt.Println("WARN: Secret key is not set- using default key. For production use 'export SECRET_KEY=...'")
		cfg.Auth.SaltKey = DefaultSecretKey
	}

	// security reasons
	_ = os.Unsetenv("SECRET_KEY")

	return &cfg, nil
}
