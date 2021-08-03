package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt"
)

// Key generates and validates JSON web tokens.
type Key struct {
	privateKey *rsa.PrivateKey
}

// GenerateRandomKey generates a random Key.
func GenerateRandomKey() (*Key, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return &Key{
		privateKey,
	}, nil
}

// ClaimsOptions are options for creating a JSON web token.
type ClaimsOptions struct {
	Lifespan time.Duration
	Id       string
}

// NewClaims generates new claims for a JSON web token.
func (k Key) NewClaims(opts ClaimsOptions) *Claims {
	return &Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(opts.Lifespan).Unix(),
			Id:        opts.Id,
			IssuedAt:  time.Now().Unix(),
		},
	}
}

// NewToken creates a new JSON web token with claims.
func (k Key) NewToken(opts ClaimsOptions) *Token {
	claims := k.NewClaims(opts)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return &Token{
		token:       token,
		keyProvider: k,
	}
}

// SignToken signs a JSON web tokens.
func (k Key) SignToken(token *jwt.Token) (string, error) {
	return token.SignedString(k.privateKey)
}

// GetVerificationKey gets the JSON web token verification key.
func (k Key) GetVerificationKey(token *jwt.Token) (interface{}, error) {
	return &k.privateKey.PublicKey, nil
}
