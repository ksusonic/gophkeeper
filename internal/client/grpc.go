package client

import (
	servicepb "github.com/ksusonic/gophkeeper/proto/service"

	"google.golang.org/grpc"
)

type GrpcClient struct {
	servicepb.AuthServiceClient
	servicepb.SecretServiceClient
}

func NewGrpcClient(address, tlsPath string) (*GrpcClient, error) {
	tlsConfig, err := loadTLSCredentials(tlsPath)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(tlsConfig))
	if err != nil {
		return nil, err
	}

	return &GrpcClient{
		servicepb.NewAuthServiceClient(conn),
		servicepb.NewSecretServiceClient(conn),
	}, nil
}
