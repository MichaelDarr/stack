package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/MichaelDarr/stack/auth/internal/config"
	pb "github.com/MichaelDarr/stack/auth/proto/auth"
	"google.golang.org/grpc"
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
	return &pb.Thing{Name: fmt.Sprintf("%d %d", some.X, some.Y)}, nil
}
