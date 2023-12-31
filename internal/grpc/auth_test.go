package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestAuthControllerGrpc_RegisterService(t *testing.T) {
	srv := grpc.NewServer()
	a := &AuthControllerGrpc{}
	a.RegisterService(srv)

	_, ok := srv.GetServiceInfo()["service.AuthService"]
	assert.True(t, ok, "server not found in registered")
}
