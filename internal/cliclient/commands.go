package cliclient

import (
	"fmt"
	"log"
	"os"

	"github.com/ksusonic/gophkeeper/internal/client"
	clipb "github.com/ksusonic/gophkeeper/proto/cli"
	"github.com/ksusonic/gophkeeper/proto/service"
	"github.com/tcnksm/go-input"
	"github.com/urfave/cli/v2"
)

const ProductServiceName = "GophKeeper"

type Helper struct {
	out          *log.Logger
	serverClient *client.GrpcClient
	storage      *Storage
	ui           *input.UI
}

func NewHelper(logger *log.Logger, serverClient *client.GrpcClient, storage *Storage) *Helper {
	return &Helper{
		serverClient: serverClient,
		storage:      storage,
		out:          logger,
		ui: &input.UI{
			Writer: logger.Writer(),
			Reader: os.Stdin,
		},
	}
}

func (h *Helper) Init(ctx *cli.Context) error {
	if h.storage.GetValue().GetToken() == "" {
		if h.askYesNo(fmt.Sprintf("Do you have an account on %s? (y/n)", ProductServiceName)) {
			return h.Login(ctx)
		} else {
			return h.Register(ctx)
		}
	}

	if h.askYesNo("Seems like you're already logged in. Logout? (y/n)") {
		h.storage.SetValue(nil)
		return h.Init(ctx)
	}
	return nil
}

func (h *Helper) Register(cCtx *cli.Context) error {
	username := h.askData("Enter your username")

	var password string
	for password == "" {
		pass := h.askPrivate("Enter your password")
		repeatedPass := h.askPrivate("Repeat your password")
		if pass != repeatedPass {
			h.out.Println("Passwords does not match")
		} else {
			password = pass
		}
	}

	register, err := h.serverClient.Register(cCtx.Context, &service.RegisterRequest{
		Login:    username,
		Password: password,
	})
	if err != nil {
		return err
	}
	fmt.Println("Got registered user:", register.String())
	h.storage.SetValue(
		&clipb.Storage{
			Token: register.AccessToken,
		},
	)
	return nil
}

func (h *Helper) Login(cCtx *cli.Context) error {
	username := h.askData("Enter your username")
	password := h.askPrivate("Enter your password")

	login, err := h.serverClient.Login(cCtx.Context, &service.LoginRequest{
		Login:    username,
		Password: password,
	})
	if err != nil {
		h.out.Printf("error at login: %v", err)
		return h.Login(cCtx)
	}
	h.storage.SetValue(
		&clipb.Storage{
			Token: login.AccessToken,
		},
	)
	return nil
}
