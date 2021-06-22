package user

import (
	"github.com/MichaelDarr/shelf/backend/internal/graphql/model"
	"github.com/google/uuid"
)

// User defines a single user.
type User struct {
	ID    uuid.UUID
	Email string
}

// GQL returns user data in the format expected by GraphQL endpoints.
func (u User) GQL() *model.User {
	return &model.User{
		ID:    u.ID,
		Email: u.Email,
	}
}
