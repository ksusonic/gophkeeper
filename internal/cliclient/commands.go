package cliclient

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ksusonic/gophkeeper/internal/client"
	datapb "github.com/ksusonic/gophkeeper/proto/data"
	servicepb "github.com/ksusonic/gophkeeper/proto/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"

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

func (h *Helper) Init(cCtx *cli.Context) error {
	if h.storage.GetToken() == "" {
		if h.askYesNo(fmt.Sprintf("Do you have an account on %s? (y/n)", ProductServiceName)) {
			return h.Login(cCtx)
		} else {
			return h.Register(cCtx)
		}
	}

	if h.askYesNo("Seems like you're already logged in. Logout? (y/n)") {
		h.storage.Token = ""
		return h.Init(cCtx)
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

	register, err := h.serverClient.Register(cCtx.Context, &servicepb.RegisterRequest{
		Login:    username,
		Password: password,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Successfully registered, %s!\n", username)
	h.storage.Token = register.AccessToken
	return nil
}

func (h *Helper) Login(cCtx *cli.Context) error {
	username := h.askData("Enter your username")
	password := h.askPrivate("Enter your password")

	login, err := h.serverClient.Login(cCtx.Context, &servicepb.LoginRequest{
		Login:    username,
		Password: password,
	})
	if err != nil {
		h.out.Printf("error at login: %v", err)
		return h.Login(cCtx)
	}
	h.storage.Token = login.AccessToken
	fmt.Printf("Successfully logged in as %s\n", username)
	return nil
}

func (h *Helper) addSecret(cCtx *cli.Context, data *datapb.Secret_Data) error {
	name := h.askData("Enter the name of secret (must be unique)")
	metadata, err := structpb.NewStruct(h.askJSON("Enter metadata in JSON format, keys are strings. If no need - leave empty."))
	if err != nil {
		return fmt.Errorf("cannot process metadata: %w", err)
	}

	ctx, cancel := context.WithTimeout(cCtx.Context, time.Second*1)
	defer cancel()

	_, err = h.serverClient.SetSecret(ctx, &servicepb.SetSecretRequest{
		Secret: &datapb.Secret{
			Name:       name,
			Meta:       metadata,
			SecretData: data,
		},
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok && s.Code() == codes.Unauthenticated {
			h.storage.Token = ""
			return fmt.Errorf("auth token was cleared: %s", s.Message())
		}
		return err
	}
	h.out.Printf("Secret %s successfully added\n", name)
	return nil
}

func (h *Helper) AddPassword(cCtx *cli.Context) error {
	login := h.askData("Enter login/email from external service account")
	password := h.askPrivate("Enter password")
	data := &datapb.Secret_Data{Variant: &datapb.Secret_Data_Authentication{Authentication: &datapb.AuthenticationData{
		Login:       login,
		RawPassword: password,
	}}}
	return h.addSecret(cCtx, data)
}

func (h *Helper) AddText(cCtx *cli.Context) error {
	text := h.askData("Enter text")
	return h.addSecret(
		cCtx,
		&datapb.Secret_Data{Variant: &datapb.Secret_Data_Text{Text: text}},
	)
}

func (h *Helper) AddBytes(cCtx *cli.Context) error {
	bytes := h.askData("Enter any byte value")
	return h.addSecret(
		cCtx,
		&datapb.Secret_Data{Variant: &datapb.Secret_Data_Any{Any: &anypb.Any{Value: []byte(bytes)}}},
	)
}

func (h *Helper) AddCard(cCtx *cli.Context) error {
	pan := h.askData("Enter your card number")
	chName := h.askData("Enter card owner name")
	expirationData := h.askData("Enter expiration data")

	return h.addSecret(
		cCtx,
		&datapb.Secret_Data{Variant: &datapb.Secret_Data_CreditCardData{CreditCardData: &datapb.CreditCardData{
			Pan:            pan,
			ChName:         chName,
			ExpirationDate: expirationData,
		}}},
	)
}
