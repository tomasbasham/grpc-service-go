package http

import (
	"net/http"

	"go.opencensus.io/plugin/ochttp"
)

// NewServer returns a HTTP server that handles tracing headers and other
// context specific configuration options.
func NewServer(handler http.Handler) (*http.Server, error) {
	if handler == nil {
		return &http.Server{Handler: http.NotFoundHandler}
	}

	return &http.Server{
		Handler: &ochttp.Handler{
			Handler: handler,
		},
	}
}
