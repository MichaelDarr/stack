package resolvers

import (
	"context"

	"github.com/MichaelDarr/shelf/backend/internal/graphql/model"
	uss "github.com/MichaelDarr/shelf/backend/internal/user/service"
)

// UserCreate creates a user.
func (r *mutationResolver) UserCreate(ctx context.Context, input model.UserCreateInput) (*model.User, error) {
	return r.UserService.CreateUser(uss.CreateUserOptions{
		Email: input.Email,
	})
}

// User fetches a single user.
func (r *queryResolver) User(ctx context.Context, input model.UserInput) (*model.User, error) {
	return r.UserService.RetrieveUserByID(input.ID)
}
