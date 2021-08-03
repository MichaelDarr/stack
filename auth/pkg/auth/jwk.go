package auth

// JWK is a JSON web key.
type JWK struct {
	// Algorithm is the algorithm intended for use with the key.
	Algorithm string `json:"alg"`
	// KeyType is the cryptographic algorithm family used with the key.
	KeyType string `json:"kty"`
	// KeyID matches a specific key.
	KeyID string `json:"kid"`
}
