package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ksusonic/gophkeeper/internal/cliclient"
	"github.com/ksusonic/gophkeeper/internal/config"
	"github.com/urfave/cli/v2"
)

var (
	Version = "dev"
	Date    = "unknown"
)

func main() {
	cfg, err := config.NewClientConfigWithStorage()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}
	storage, err := cliclient.NewStorage(cfg.StoragePath, !cfg.Debug)
	if err != nil {
		log.Fatalf("could not load storage: %v", err)
	}

	defer func(cfg *config.ClientConfig) {
		err := storage.Save()
		if err != nil {
			log.Printf("Sorry, could not save storage: %v\n", err)
		}
	}(cfg)

	app := &cli.App{
		EnableBashCompletion: true,
		Usage:                "Keeps your secrets in the air!",
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "get version",
				Action: func(cCtx *cli.Context) error {
					fmt.Printf("version: '%s' from %s\n", Version, Date)
					return nil
				},
			},
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "register or login to system",
				Action: func(cCtx *cli.Context) error {
					// TODO register or login to system
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a", "ad"},
				Usage:   "add new secret",
				Subcommands: []*cli.Command{
					{
						Name:  "password",
						Usage: "authentication data",
						Action: func(cCtx *cli.Context) error {
							return nil
						},
					},
					{
						Name:  "text",
						Usage: "text value",
						Action: func(cCtx *cli.Context) error {
							return nil
						},
					},
					{
						Name:  "byte",
						Usage: "byte data",
						Action: func(cCtx *cli.Context) error {
							return nil
						},
					},
					{
						Name:  "card",
						Usage: "credit card data",
						Action: func(cCtx *cli.Context) error {
							return nil
						},
					},
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"d", "delete"},
				Usage:   "delete secret",
				Action: func(context *cli.Context) error {
					// TODO delete secret
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
