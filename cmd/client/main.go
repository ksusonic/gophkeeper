package main

import (
	"log"
	"os"

	"github.com/ksusonic/gophkeeper/internal/client"
	"github.com/ksusonic/gophkeeper/internal/clihelper"
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
	storage, err := clihelper.NewStorage(cfg.StoragePath, !cfg.Debug)
	if err != nil {
		out.Fatalf("could not load storage: %v", err)
	}
	grpc, err := client.NewGrpcClient(cfg.ServerURL, cfg.CertPath)
	if err != nil {
		log.Fatalf("error creating grpc client: %v", err)
	}
	cliClient := clihelper.NewHelper(out, grpc, storage)

	cli.VersionPrinter = func(cCtx *cli.Context) {
		out.Printf("version: '%s' from %s\n", cCtx.App.Version, Date)
	}
	initCommand := &cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "register or login to system",
		Action:  cliClient.Init,
	}
	app := &cli.App{
		Name:                 "gophkeeper",
		Version:              Version,
		Usage:                "Keeps your secrets in the air!",
		EnableBashCompletion: true,

		Action: func(ctx *cli.Context) error {
			if !storage.TokenIsValid(ctx.Context) {
				out.Println("Welcome to GophKeeper!")
				return initCommand.Run(ctx)
			}
			out.Println("You are logged in")
			return cli.ShowAppHelp(ctx)
		},
		Commands: cli.Commands{
			initCommand,
			{
				Name:    "add",
				Aliases: []string{"a", "ad"},
				Usage:   "add new secret",
				Before:  storage.LoginInterceptor,
				Subcommands: []*cli.Command{
					{
						Name:   "password",
						Usage:  "authentication data",
						Action: cliClient.AddPassword,
					},
					{
						Name:   "text",
						Usage:  "text value",
						Action: cliClient.AddText,
					},
					{
						Name:   "byte",
						Usage:  "byte data",
						Action: cliClient.AddBytes,
					},
					{
						Name:   "card",
						Usage:  "credit card data",
						Action: cliClient.AddCard,
					},
				},
			},
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "get secret",
				Before:  storage.LoginInterceptor,
				Action:  cliClient.GetSecret,
				Subcommands: []*cli.Command{
					{
						Name:    "all",
						Aliases: []string{"a"},
						Usage:   "get all secret",
						Action:  cliClient.GetAllSecrets,
					},
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"d", "delete"},
				Usage:   "delete secret",
				Before:  storage.LoginInterceptor,
				Action:  cliClient.RemoveSecret,
			},
		},
	}
	defer func(storage *clihelper.Storage) {
		if err := storage.Save(); err != nil {
			out.Fatalf("cannot save storage: %v", err)
		}
	}(storage)

	if err := app.Run(os.Args); err != nil {
		out.Fatal(err)
	}
}
