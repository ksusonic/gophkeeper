package main

import (
	"log"
	"os"

	"github.com/ksusonic/gophkeeper/internal/cliclient"
	"github.com/ksusonic/gophkeeper/internal/client"
	"github.com/ksusonic/gophkeeper/internal/config"
	"github.com/urfave/cli/v2"
)

var (
	Version = "dev"
	Date    = "unknown"

	out = log.New(os.Stderr, "", 0)
)

func main() {
	cfg, err := config.NewClientConfigWithStorage()
	if err != nil {
		out.Fatalf("could not load config: %v", err)
	}
	storage, err := cliclient.NewStorage(cfg.StoragePath, !cfg.Debug)
	if err != nil {
		out.Fatalf("could not load storage: %v", err)
	}

	defer func(cfg *config.ClientConfig) {
		err := storage.Save()
		if err != nil {
			out.Printf("Sorry, could not save storage: %v\n", err)
		}
	}(cfg)

	grpc, err := client.NewGrpcClient(cfg.ServerURL, cfg.CertPath)
	if err != nil {
		log.Fatalf("error creating grpc client: %v", err)
	}

	cli.VersionPrinter = func(cCtx *cli.Context) {
		out.Printf("version: '%s' from %s\n", cCtx.App.Version, Date)
	}

	initCommand := &cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "register or login to system",
		Action: func(ctx *cli.Context) error {
			return cliclient.NewHelper(out, grpc, storage).Init(ctx)
		},
	}
	app := &cli.App{
		Name:    "gophkeeper",
		Version: Version,
		Usage:   "Keeps your secrets in the air!",

		EnableBashCompletion: true,
		Commands: cli.Commands{
			initCommand,
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
		Action: func(ctx *cli.Context) error {
			if storage.GetToken() == "" {
				out.Println("Welcome to GophKeeper!")
				return initCommand.Run(ctx)
			}
			out.Println("You are logged in")
			// TODO: check if token is valid
			return cli.ShowAppHelp(ctx)
		},
	}

	if err := app.Run(os.Args); err != nil {
		out.Fatal(err)
	}
}
