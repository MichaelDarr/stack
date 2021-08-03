package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"

	"github.com/golang-jwt/jwt"
)

// RSAKey is an RSA key.
type RSAKey struct {
	privateKey *rsa.PrivateKey
}

// RSAJWK defines additional JWK fields included in RSA JSON web keys.
type RSAJWK struct {
	BaseJWK
	// Modulus is the RSA key's modulus value.
	Modulus string `json:"n"`
	// Exponent is the RSA key's exponent value.
	Exponent string `json:"e"`
	// X5T is the X.509 certificate SHA-1 thumbprint.
	X5T string `json:"x5t"`
}

// GenerateRandomRSAKey generates a random RSA JSON web key.
func GenerateRandomRSAKey() (*RSAKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return &RSAKey{
		privateKey,
	}, nil
}

// GetJWK gets JSON web key fields.
func (r RSAKey) GetJWK() interface{} {
	return RSAJWK{
		BaseJWK: BaseJWK{
			Algorithm: "RS256",
			KeyType:   "RSA",
			KeyID:     r.GetKeyID(),
		},
		Modulus:  r.GetModulus(),
		Exponent: r.GetExponent(),
		X5T:      r.GetX5T(),
	}
}

// GetExponent gets the Base64urlUInt-encoded RSA exponent value.
func (r RSAKey) GetExponent() string {
	exp := int64(r.GetPublicKey().E)
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(buf, exp)
	encoded := base64.URLEncoding.EncodeToString(buf)
	return encoded
}

// GetKeyID gets the key identifier.
func (r RSAKey) GetKeyID() string {
	return r.GetX5T()
}

// GetModulus gets the Base64urlUInt-encoded RSA modulus value.
func (r RSAKey) GetModulus() string {
	modBytes := r.GetPublicKey().N.Bytes()
	encoded := base64.URLEncoding.EncodeToString(modBytes)
	return encoded
}

// GetPublicKey gets the public RSA key.
func (r RSAKey) GetPublicKey() *rsa.PublicKey {
	return &r.privateKey.PublicKey
}

// GetSigningMethod gets the RSA signing method.
func (r RSAKey) GetSigningMethod() jwt.SigningMethod {
	return jwt.SigningMethodRS256
}

// GetVerificationKey gets the JSON web token verification key.
func (r RSAKey) GetVerificationKey() interface{} {
	return &r.privateKey.PublicKey
}

// GetX5T gets the X.509 certificate SHA-1 thumbprint.
// https://datatracker.ietf.org/doc/html/rfc7517#section-4.8
func (r RSAKey) GetX5T() string {
	der := x509.MarshalPKCS1PublicKey(r.GetPublicKey())
	sum := sha1.Sum(der)
	x5t := base64.URLEncoding.EncodeToString(sum[:])
	return x5t
}

// SignToken signs a JSON web tokens.
func (r RSAKey) SignToken(token *jwt.Token) (string, error) {
	return token.SignedString(r.privateKey)
}
