package option

import (
	"github.com/tomasbasham/grpc-service-go/internal"
	"google.golang.org/grpc"
)

// ClientOption is an option for an API client.
type ClientOption interface {
	Apply(*internal.ClientOptions)
}

type endpoint string

// WithEndpoint creates an option from the given endpoint.
func WithEndpoint(e string) ClientOption {
	return endpoint(e)
}

// Apply sets the endpoint on the configuration value type.
func (e endpoint) Apply(cfg *internal.ClientOptions) {
	cfg.Endpoint = string(e)
}

// Wrapped in a struct because grpc.DialOption is an interface.
type gRPCDialOption struct{ grpc.DialOption }

// WithGRPCDialOption creates an option from a gRPC dial option.
func WithGRPCDialOption(o grpc.DialOption) ClientOption {
	return gRPCDialOption{o}
}

// Apply appends a gRPC dial option to a slace of options on the configuration
// value type.
func (o gRPCDialOption) Apply(cfg *internal.ClientOptions) {
	cfg.GRPCDialOptions = append(cfg.GRPCDialOptions, o.DialOption)
}
