package config

import (
	"fmt"
	"os"

	clipb "github.com/ksusonic/gophkeeper/proto/cli"
	"google.golang.org/protobuf/proto"

	"github.com/caarlos0/env/v8"
)

type ClientConfigWithStorage struct {
	ServerURL string `env:"SERVER" envDefault:":3000"`
	CertPath  string `env:"CERT" envDefault:"cert/server-tls/ca-cert.pem"`
	Debug     bool   `env:"DEBUG" envDefault:"false"`

	storagePath string `env:"STORAGE" envDefault:"/tmp/gophkeeper-storage"`
	storage     *clipb.Storage
}

func (s *ClientConfigWithStorage) Save() error {
	marshal, err := proto.Marshal(s.storage)
	if err != nil {
		return err
	}
	err = os.WriteFile(s.storagePath, marshal, 0666)
	if err != nil {
		return err
	}
	return nil
}

func NewClientConfigWithStorage() (*ClientConfigWithStorage, error) {
	var cfg ClientConfigWithStorage
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	file, err := os.ReadFile(cfg.storagePath)
	if file != nil {
		storage := &clipb.Storage{}
		err = proto.Unmarshal(file, storage)
		if err != nil {
			fmt.Printf("Could not unmarshall proto: %v, ignoring previous data\n", err)
		} else {
			cfg.storage = storage
		}
	} else if cfg.Debug {
		fmt.Printf("Could not load token file: %v\n", err)
	}

	return &cfg, nil
}
