package server

import (
	"context"
	"time"

	"github.com/MichaelDarr/stack/auth/pkg/auth"
	pb "github.com/MichaelDarr/stack/auth/proto/auth"
)

// AuthServer is an authentication GRPC server.
type AuthServer struct {
	pb.UnimplementedAuthServer
	keySet *auth.KeySet
}

// GetToken gets an authentication token.
func (g *AuthServer) GetToken(ctx context.Context, req *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	token, err := g.keySet.NewToken(auth.ClaimsOptions{
		Lifespan: time.Hour * 12,
		Id:       req.Id,
	})
	if err != nil {
		return nil, err
	}
	signed, err := token.Sign()
	if err != nil {
		return nil, err
	}
	return &pb.GetTokenResponse{Token: signed}, nil
}

// ValidateToken validates an authentication token.
func (g *AuthServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	_, err := g.keySet.ParseToken(req.Token)
	return &pb.ValidateTokenResponse{Valid: err == nil}, nil
}
