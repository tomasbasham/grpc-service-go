package credentials

import "context"

// PerRequestCredentials defines a common interface for the HTTP credentials
// that will be attached to each request as security information.
type PerRequestCredentials interface {
	// Credentials returns the next valid request credentails. This should only
	// be invoked by the transport mechanism, most likely through a
	// http.RoundTripper.
	Credentials(ctx context.Context) (map[string]string, error)
}
