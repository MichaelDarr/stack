package auth

import "github.com/golang-jwt/jwt"

// Claims are JSON web token claims.
type Claims struct {
	jwt.StandardClaims
}
