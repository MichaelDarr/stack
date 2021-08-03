package server

import (
	"fmt"
	"net"

	"github.com/MichaelDarr/stack/auth/pkg/auth"
	pb "github.com/MichaelDarr/stack/auth/proto/auth"
	"google.golang.org/grpc"
)

// GRPC is the a GRPC server.
type GRPC struct {
	grpcServer *grpc.Server
	authServer *AuthServer
}

// NewGRPC creates an auth grpc server.
func NewGRPC(keySet *auth.KeySet) *GRPC {
	grpcServer := grpc.NewServer()
	authServer := &AuthServer{keySet: keySet}
	pb.RegisterAuthServer(grpcServer, authServer)

	return &GRPC{
		grpcServer,
		authServer,
	}
}

// Serve accepts incoming connections.
func (g GRPC) Serve(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	return g.grpcServer.Serve(lis)
}
