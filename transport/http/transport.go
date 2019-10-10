package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tomasbasham/grpc-service-go/transport/http/credentials"
)

type credentialsKey struct{}

const Credentials = credentialsKey{}

type perRequestCredentialTransport struct {
	creds credentials.PerRequestCredentials
}

func WithPerRequestCredentialTransport(creds credentials.PerRequestCredentials) http.RoundTripper {
	return &perRequestCredentialTransport{creds: creds}
}

func (t *perRequestCredentialTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	m, err := t.creds.Credentials(r.Context())
	if err != nil {
		return nil, fmt.Errorf("", err)
	}

	ctx := context.WithValue(r.Context(), Credentials, m)
	req := r.WithContext(ctx)
	// Somehow call the next trnasport and return
}
