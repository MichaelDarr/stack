package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

// Key is a JSON web key.
type Key interface {
	// GetSigningMethod gets the signing method.
	GetSigningMethod() jwt.SigningMethod
	// GetJWK gets the JSON web key which will be marshalled to JSON.
	GetJWK() interface{}
	// GetKeyID gets the key identifier.
	GetKeyID() string
	// GetVerificationKey gets the JSON web token verification key.
	GetVerificationKey() interface{}
	// Sign signs a JSON web tokens.
	SignToken(token *jwt.Token) (string, error)
}

// KeySet is a JSON web key set.
type KeySet struct {
	keys []Key
}

// NewKeySet creates a new JSON web key set.
func NewKeySet(keys []Key) *KeySet {
	return &KeySet{keys: keys}
}

// GetVerificationKey gets the JSON web token verification key.
func (ks KeySet) GetVerificationKey(token *jwt.Token) (interface{}, error) {
	tokenKeyIDHeader, ok := token.Header["kid"]
	if !ok {
		return nil, errors.New("token header 'kid' missing")
	}
	tokenKeyID := fmt.Sprintf("%v", tokenKeyIDHeader)
	for _, key := range ks.keys {
		if tokenKeyID == key.GetKeyID() {
			return key.GetVerificationKey(), nil
		}
	}
	return nil, fmt.Errorf("token key id not found in set: %v", tokenKeyID)
}

// NewToken creates a new JSON web token with claims.
func (ks KeySet) NewToken(claimOpts ClaimsOptions) (*Token, error) {
	if len(ks.keys) == 0 {
		return nil, errors.New("token set empty")
	}
	key := ks.keys[0]
	return NewToken(key, claimOpts), nil
}

// ParseToken parses and validates a JSON web token.
func (keySet KeySet) ParseToken(tokenString string) (*Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, keySet.GetVerificationKey)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return &Token{token: token}, nil
}
