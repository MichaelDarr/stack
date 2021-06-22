package user

import (
	"github.com/MichaelDarr/shelf/backend/internal/database/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRepository interacts with user database information.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

// CreateUserOptions are options for creating a user.
type CreateUserOptions struct {
	Email string
}

// CreateUser creates a new user.
func (s *UserRepository) CreateUser(opts CreateUserOptions) (models.User, error) {
	user := models.User{
		Email: opts.Email,
	}
	result := s.db.Create(&user)
	return user, result.Error
}

// RetrieveUserByID retrieves a user by ID.
func (s *UserRepository) RetrieveUserByID(id uuid.UUID) (models.User, error) {
	var user models.User
	result := s.db.First(&user, id)
	return user, result.Error
}
