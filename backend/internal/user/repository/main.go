package user

import (
	"github.com/MichaelDarr/shelf/backend/internal/database/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRepository interacts with user database information
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

// CreateUser creates a new user
func (s *UserRepository) CreateUser(name string) (models.User, error) {
	user := models.User{
		Name: name,
	}
	result := s.db.Create(&user)
	return user, result.Error
}

// RetrieveUserByID retrieves a user by ID
func (s *UserRepository) RetrieveUserByID(id uuid.UUID) (models.User, error) {
	var user models.User
	result := s.db.First(&user, id)
	return user, result.Error
}
