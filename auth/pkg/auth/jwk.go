package auth

// BaseJWK defines JSON web key fields returned by all key types.
type BaseJWK struct {
	// Algorithm is the algorithm intended for use with the key.
	Algorithm string `json:"alg"`
	// KeyType is the cryptographic algorithm family used with the key.
	KeyType string `json:"kty"`
	// KeyID matches a specific key.
	KeyID string `json:"kid"`
}

// JWKS is a JSON web key set.
type JWKS struct {
	Keys []interface{} `json:"keys"`
}
