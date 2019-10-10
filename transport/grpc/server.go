package grpc

import (
	"github.com/tomasbasham/grpc-service-go/internal"
	"github.com/tomasbasham/grpc-service-go/option"

	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
)

// RegisterServer returns a gRPC server that handles tracing headers and other
// context specific configuration options.
func RegisterServer(opts ...option.ServerOption) (*grpc.Server, error) {
	var cfg internal.ServerOptions
	for _, opt := range opts {
		opt.Apply(&cfg)
	}

	if err := cfg.Valid(); err != nil {
		return nil, err
	}

	grpcOpts := []grpc.ServerOption{
		grpc.StatsHandler(&ocgrpc.ServerHandler{}),
	}

	grpcOpts = append(grpcOpts, cfg.GRPCServerOptions...)
	return grpc.NewServer(grpcOpts...)
}
