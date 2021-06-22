package resolvers

import (
	"github.com/MichaelDarr/shelf/backend/internal/config"
	"github.com/MichaelDarr/shelf/backend/internal/graphql/generated"
	"github.com/MichaelDarr/shelf/backend/internal/user"
)

type Resolver struct {
	Config      *config.ServerConfig
	UserService *user.Service
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
