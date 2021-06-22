package resolvers

import (
	"context"

	"github.com/MichaelDarr/shelf/backend/internal/graphql/model"
	"github.com/MichaelDarr/shelf/backend/internal/user"
)

// UserCreate creates a user.
func (r *mutationResolver) UserCreate(ctx context.Context, input model.UserCreateInput) (*model.User, error) {
	user, err := r.UserService.Create(user.CreateUserOptions{
		Email: input.Email,
	})
	if err != nil {
		return nil, err
	}
	return user.GQL(), nil
}

// User fetches a single user.
func (r *queryResolver) User(ctx context.Context, input model.UserInput) (*model.User, error) {
	user, err := r.UserService.Get(input.ID)
	if err != nil {
		return nil, err
	}
	return user.GQL(), nil
}
