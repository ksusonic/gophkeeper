package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestSecretControllerGrpc_RegisterService(t *testing.T) {
	srv := grpc.NewServer()
	a := &SecretControllerGrpc{}
	a.RegisterService(srv)

	_, ok := srv.GetServiceInfo()["service.SecretService"]
	assert.True(t, ok, "server not found in registered")
}
