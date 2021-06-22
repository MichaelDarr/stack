package user

import (
	"github.com/MichaelDarr/shelf/backend/internal/graphql/model"
	r "github.com/MichaelDarr/shelf/backend/internal/user/repository"
	"github.com/google/uuid"
)

// UserService is a configured database UserService
type UserService struct {
	repo *r.UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo *r.UserRepository) *UserService {
	return &UserService{repo}
}

// CreateUserOptions are options for creating a user.
type CreateUserOptions struct {
	Email string
}

// CreateUser creates a user
func (s *UserService) CreateUser(opts CreateUserOptions) (*model.User, error) {
	user, err := s.repo.CreateUser(r.CreateUserOptions{
		Email: opts.Email,
	})
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

// RetrieveUserByID retrieves a user by ID
func (s *UserService) RetrieveUserByID(id uuid.UUID) (*model.User, error) {
	user, err := s.repo.RetrieveUserByID(id)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   user.ID,
		Email: user.Email,
	}, nil
}
