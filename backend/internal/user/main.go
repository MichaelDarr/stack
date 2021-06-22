package user

import "github.com/google/uuid"

// Respository provides access to user data.
type Respository interface {
	Create(opts CreateUserOptions) (*User, error)
	Get(id uuid.UUID) (*User, error)
}

// Service is the user service.
type Service struct {
	repo Respository
}

// NewService creates a user service.
func NewService(repo Respository) *Service {
	return &Service{repo}
}

// Create creates a user.
func (s Service) Create(opts CreateUserOptions) (*User, error) {
	return s.repo.Create(opts)
}

// Get retrieves a user by their ID.
func (s Service) Get(id uuid.UUID) (*User, error) {
	return s.repo.Get(id)
}
