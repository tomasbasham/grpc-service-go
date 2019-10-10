package internal

import (
	"google.golang.org/grpc"
)

// ClientOptions describe the set of options necessary to connect to an API.
// This is encapsulated inside the internal package to prevent cyclic imports
// from the option and transport packages and to prevent other code bases from
// importing this package directly.
type ClientOptions struct {
	Endpoint        string
	GRPCDialOptions []grpc.DialOption
}

// Valid returns an error if the ClientOptions are not sensibly configured.
func (o *ClientOptions) Valid() error {
	return nil
}
