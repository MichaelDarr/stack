package resolvers

import (
	"context"
	"fmt"

	"github.com/MichaelDarr/shelf/backend/internal/graphql/model"
)

func (r *mutationResolver) UserCreate(ctx context.Context, name string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	if id != "a42c9822-f24b-40ce-9416-eecf2eb9d487" {
		return nil, fmt.Errorf("user id not found: %s", id)
	}
	return &model.User{
		ID:   "a42c9822-f24b-40ce-9416-eecf2eb9d487",
		Name: "Michael Darr",
	}, nil
}
