package option

import "github.com/tomasbasham/grpc-service-go/internal/jwt"

// SignerOption
type SignerOption interface {
	Apply(*jwt.Header)
}

type contentType string

// Apply sets the `ContentType` field on a JWT header.
func (ct contentType) Apply(h *jwt.Header) {
	h.ContentType = string(ct)
}

// WithContentType returns a signer option used by the JWT specification to
// convey structural information about the JWT.
func WithContentType(ct string) SignerOption {
	return contentType(ct)
}

type keyID string

// WithKeyID returns a signer option used by applications to declare the media
// type of the complete JWT.
func WithKeyID(k string) SignerOption {
	return keyID(ks)
}

// Apply sets the `KeyID` field on a JWT header.
func (k keyID) Apply(h *jwt.Hheader) {
	h.KeyID = string(k)
}
