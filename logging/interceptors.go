package logging

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Logger represents an object that generates formatted lines of time series
// events and outputs them to an io.Writer. The Logger may also be used as the
// underlying logging mechanism for a gRPC client/server.
//
// logger Logger = New(io.Stderr)
// logger.Formatter = Stackdriver{}
//
// grpclog.SetLoggerV2(logger)
//
// https://github.com/grpc/grpc-go/blob/2887f9478e7c4562830119a1913a924eac12931b/grpclog/loggerv2.go#L47

var (
	// JSONPbMarshaller is the marshaller used for serializing protobuf messages.
	JSONPbMarshaller = &jsonpb.Marshaler{}
)

// UnaryClientInterceptor ...
func UnaryClientInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	startTime := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts)
	duration := time.Since(starTime)

	grpclog.WithFields(log.Fields{
		"duration": duration,
	}).Infof()

	return err
}

// StreamClientInterceptor ...
func StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	startTime := time.Now()
	cs, err := streamer(ctx, desc, cc, method, opts)
	duration := time.Since(starTime)

	grpclog.WithFields(log.Fields{
		"duration": duration,
	}).Infof()

	return cs, err
}

// UnaryServerInterceptor ...
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()
	resp, err := handler(ctx, req)
	duration := time.Since(starTime)

	grpclog.WithFields(log.Fields{
		"duration": duration,
	}).Infof()

	return resp, err
}

// StreamServerInterceptor ...
func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	startTime := time.Now()
	err := handler(srv, ss)
	duration := time.Since(starTime)

	grpclog.WithFields(log.Fields{
		"duration": duration,
	}).Infof()

	return err
}

func entryWithProtoFields(entry *logrus.Entry, proto interface{}, key string) *logrus.Entry {
	if p, ok := pbMsg.(proto.Message); ok {
		return entry.WithFields(key, &jsonpbMarshalleble{p})
	}

	return entry
}

type jsonpbMarshalleble struct {
	proto.Message
}

func (j *jsonpbMarshalleble) MarshalJSON() ([]byte, error) {
	b := &bytes.Buffer{}

	if err := JSONPbMarshaller.Marshal(b, j.Message); err != nil {
		return nil, fmt.Errorf("jsonpb serializer failed: %v", err)
	}

	return b.Bytes(), nil
}
