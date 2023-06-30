package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	servicepb "github.com/ksusonic/gophkeeper/proto/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

func loadTLSCredentials(tlsPath string) (credentials.TransportCredentials, error) {
	pemServerCA, err := os.ReadFile(tlsPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(tlsConfig), nil
}
