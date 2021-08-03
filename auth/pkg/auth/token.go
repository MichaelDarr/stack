package auth

import (
	"github.com/golang-jwt/jwt"
)

// Token is a JSON web token.
type Token struct {
	token *jwt.Token
	key   Key
}

// NewToken creates a new JSON web token with claims.
func NewToken(key Key, claimOpts ClaimsOptions) *Token {
	claims := NewClaims(claimOpts)
	token := jwt.NewWithClaims(key.GetSigningMethod(), claims)
	token.Header["kid"] = key.GetKeyID()
	return &Token{
		token: token,
		key:   key,
	}
}

// Sign gets the signed JSON web token string.
func (t Token) Sign() (string, error) {
	return t.key.SignToken(t.token)
}

// Claims gets a token's claims.
func (t Token) Claims() (*Claims, bool) {
	claims, ok := t.token.Claims.(*Claims)
	return claims, ok
}
