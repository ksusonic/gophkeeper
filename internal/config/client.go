package config

import (
	"fmt"
	"os"

	clipb "github.com/ksusonic/gophkeeper/proto/cli"
	"google.golang.org/protobuf/proto"

	"github.com/caarlos0/env/v8"
)

type ClientStorage struct {
	ServerURL string `env:"SERVER" envDefault:":3000"`
	TokenPath string `env:"TOKEN" envDefault:"/tmp/gophkeeper/token"`
	Debug     bool   `env:"DEBUG" envDefault:"false"`

	storage *clipb.Storage
}

func LoadClientStorage() (*ClientStorage, error) {
	var cfg ClientStorage
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	file, err := os.ReadFile(cfg.TokenPath)
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
