package internal

import (
	"context"
	"time"
)

// DurationFunc defines a type used to return a datetime representation given
// the current datetime.
type DurationFunc func(duration time.Duration) interface{}

// TraceContextFunc defines a type used to return a trace context
// representation given a context.
type TraceContextFunc func(ctx context.Context) interface{}

// LoggingOptions describes the set of options available to configure the
// logging package. This is encapsulated inside the internal package to prevent
// cyclic imports from the option package and to prevent other code bases from
// importing this package directly.
type LoggingOptions struct {
	Duration DurationFunc
	Trace    TraceContextFunc
}

// Valid returns an error if the LoggingOptions are not sensibly configured.
func (o *LoggingOptions) Valid() error {
	return nil
}
