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
	key *auth.Key
}

// GetToken gets an authentication token.
func (g *AuthServer) GetToken(ctx context.Context, req *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	token := g.key.NewToken(auth.ClaimsOptions{
		Lifespan: time.Hour * 12,
		Id:       req.Id,
	})
	signed, err := token.Sign()
	if err != nil {
		return nil, err
	}
	return &pb.GetTokenResponse{Token: signed}, nil
}

// ValidateToken validates an authentication token.
func (g *AuthServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	_, err := auth.ParseToken(req.Token, g.key)
	return &pb.ValidateTokenResponse{Valid: err == nil}, nil
}
