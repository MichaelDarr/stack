package resolvers

import (
	"context"

	"github.com/MichaelDarr/shelf/backend/internal/graphql/model"
	"github.com/google/uuid"
)

// UserCreate creates a user
func (r *mutationResolver) UserCreate(ctx context.Context, name string) (*model.User, error) {
	return r.UserService.CreateUser(name)
}

// User fetches a single user
func (r *queryResolver) User(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return r.UserService.RetrieveUserByID(id)
}
