package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/MichaelDarr/shelf/auth/proto"
	"google.golang.org/grpc"
)

type authServer struct {
	pb.UnimplementedAuthServer
}

// GetSomething gets something.
func (s *authServer) GetSomething(ctx context.Context, some *pb.Some) (*pb.Thing, error) {
	return nil, nil
}

func newServer() *authServer {
	s := &authServer{}
	return s
}

func main() {
	port := 3334
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
