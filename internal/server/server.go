package server

import (
	authpb "github.com/ksusonic/gophkeeper/proto/auth"
	servicepb "github.com/ksusonic/gophkeeper/proto/service"
)

type Server struct {
	authpb.UnimplementedAuthServiceServer
	servicepb.UnimplementedDataServiceServer
}
