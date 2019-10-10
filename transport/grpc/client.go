package grpc

import (
	"context"

	"github.com/tomasbasham/grpc-service-go/internal"
	"github.com/tomasbasham/grpc-service-go/option"

	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
)

// Dial returns a gRPC client setup for tracing and other context specific
// configuration options.
func Dial(ctx context.Context, opts ...option.ClientOption) (*grpc.ClientConn, error) {
	var cfg internal.ClientOptions
	for _, opt := range opts {
		opt.Apply(&cfg)
	}

	if err := cfg.Valid(); err != nil {
		return nil, err
	}

	grpcOpts := []grpc.DialOption{
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
	}

	grpcOpts = append(grpcOpts, cfg.GRPCDialOptions...)
	return grpc.DialContext(ctx, cfg.Endpoint, grpcOpts...)
}

// DialInsecure returns an insecure gRPC client connection setup for tracing
// and other context specific configuration options.
func DialInsecure(ctx context.Context, opts ...option.ClientOption) (*grpc.ClientConn, error) {
	return Dial(ctx, append(opts, option.WithGRPCDialOption(grpc.WithInsecure()))...)
}
