package config

import (
	"github.com/caarlos0/env/v8"
)

type ClientConfig struct {
	ServerURL   string `env:"SERVER" envDefault:":3000"`
	CertPath    string `env:"CERT" envDefault:"cert/server-tls/ca-cert.pem"`
	StoragePath string `env:"STORAGE" envDefault:"/tmp/gophkeeper-storage"`
	Debug       bool   `env:"DEBUG" envDefault:"false"`
}

func NewClientConfigWithStorage() (*ClientConfig, error) {
	var cfg ClientConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
