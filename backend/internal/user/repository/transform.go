package userRepository

import (
	"github.com/MichaelDarr/stack/backend/internal/database/models"
	"github.com/MichaelDarr/stack/backend/internal/user"
)

// dbUserToServiceUser transforms a user database record to a user compatable with the user service.
func dbUserToServiceUser(u *models.User) *user.User {
	return &user.User{
		ID:    u.ID,
		Email: u.Email,
	}
}
