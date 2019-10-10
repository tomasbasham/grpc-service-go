package jwt

// Header is a JWT header as specified in RFC 7519.
type Header struct {
	Algorithm   string `json:"alg"`           // The algorithm used for signature.
	Type        string `json:"typ"`           // Represents the token type.
	ContentType string `json:"cty,omitempty"` // Represents a hint of token content type.
	KeyID       string `json:"kid,omitempty"` // The optional hint of which key is being used.
}
