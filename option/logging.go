package option

import "github.com/tomasbasham/grpc-service-go/internal"

// LoggingOption is an option for a logger.
type LoggingOption interface {
	Apply(*internal.LoggingOptions)
}

type duration struct{ internal.DurationFunc }

// WithDuration creates an option to calculate a duration using the given
// function.
func WithDuration(f internal.DurationFunc) LoggingOption {
	return duration{f}
}

// Apply sets the duration function on the configuration value type.
func (d duration) Apply(cfg *internal.LoggingOptions) {
	cfg.Duration = d.DurationFunc
}

type trace struct{ internal.TraceContextFunc }

// WithTraceContext creates an option to create a trace context using the given
// function.
func WithTraceContext(f internal.TraceContextFunc) LoggingOption {
	return trace{f}
}

// Apply sets the trace context function on the configuration value type.
func (t trace) Apply(cfg *internal.LoggingOptions) {
	cfg.Trace = t.TraceContextFunc
}
