package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// Claims are JSON web token claims.
type Claims struct {
	jwt.StandardClaims
}

// ClaimsOptions are options for creating a JSON web token.
type ClaimsOptions struct {
	Lifespan time.Duration
	Id       string
}

// NewClaims generates new claims for a JSON web token.
func NewClaims(opts ClaimsOptions) *Claims {
	return &Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(opts.Lifespan).Unix(),
			Id:        opts.Id,
			IssuedAt:  time.Now().Unix(),
		},
	}
}
