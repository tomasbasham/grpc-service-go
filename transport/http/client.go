package http

import (
	"net/http"

	"go.opencensus.io/plugin/ochttp"
)

// Dial returns a HTTP client with setup for tracing and other context specific
// configuration options.
func Dial(transport http.RoundTripper) (*http.Client, error) {
	if transport == nil {
		transport = http.DefaultTransport
	}

	return &http.Client{
		Transport: &ochttp.Transport{
			Base: transport,
		},
	}
}
