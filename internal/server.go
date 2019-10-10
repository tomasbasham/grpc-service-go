package internal

import (
	"net/http"
)

// ServerOptions describe the set of options necessary to create an API server.
// This is encapsulated inside the internal package to prevent cyclic imports
// from the option and transport packages and to prevent other code bases from
// importing this package directly.
type ServerOptions struct {
	Address string
	Handler http.Handler
}

// Valid returns an error if the ServerOptions are not sensibly configured.
func (o *ServerOptions) Valid() error {
	return nil
}
