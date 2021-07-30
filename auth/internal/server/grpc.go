package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/MichaelDarr/shelf/auth/proto/auth"
	"github.com/MichaelDarr/shelf/auth/internal/config"
)

// GRPC starts the grpc server.
func GRPC(cfg *config.ServerConfig) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, newGrpcServer())
	grpcServer.Serve(lis)
}

type grpcServer struct {
	pb.UnimplementedAuthServer
}

func newGrpcServer() *grpcServer {
	s := &grpcServer{}
	return s
}

// GetSomething gets something.
func (g *grpcServer) GetSomething(ctx context.Context, some *pb.Some) (*pb.Thing, error) {
	return nil, nil
}
