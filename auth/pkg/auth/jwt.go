package auth

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

// KeyProvider provides JSON web token key functionality.
type KeyProvider interface {
	// Sign signs a JSON web tokens.
	SignToken(token *jwt.Token) (string, error)
	// GetVerificationKey gets the JSON web token verification key.
	GetVerificationKey(token *jwt.Token) (interface{}, error)
}

// Token is a JSON web token.
type Token struct {
	token       *jwt.Token
	keyProvider KeyProvider
}

// ParseToken parses and validates a JSON web token.
func ParseToken(tokenString string, keyProvider KeyProvider) (*Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, keyProvider.GetVerificationKey)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return &Token{token: token}, nil
}

// Claims gets a token's claims.
func (t Token) Claims() (*Claims, bool) {
	claims, ok := t.token.Claims.(*Claims)
	return claims, ok
}

// Sign gets the signed JSON web token string.
func (t Token) Sign() (string, error) {
	return t.keyProvider.SignToken(t.token)
}
