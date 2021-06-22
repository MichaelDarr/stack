package userRepository

import (
	"fmt"

	"github.com/MichaelDarr/shelf/backend/internal/database/models"
	"github.com/MichaelDarr/shelf/backend/internal/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// repository interacts with users in the database.
type repository struct {
	db *gorm.DB
}

// New creates a new user repository.
func New(db *gorm.DB) *repository {
	return &repository{db}
}

// Create creates a new user.
func (r repository) Create(opts user.CreateUserOptions) (user *user.User, err error) {
	dbUser := &models.User{
		Email: opts.Email,
	}
	result := r.db.Create(dbUser)
	if err = result.Error; err == nil {
		user = dbUserToServiceUser(dbUser)
	}
	return
}

// Get retrieves a user by ID.
func (r *repository) Get(id uuid.UUID) (user *user.User, err error) {
	var dbUser models.User
	fmt.Println(id)
	result := r.db.First(&dbUser, id)
	if err = result.Error; err == nil {
		user = dbUserToServiceUser(&dbUser)
	}
	return
}
