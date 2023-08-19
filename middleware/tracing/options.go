package tracing

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// Option is tracing option.
type Option func(*options)

type options struct {
	tracerProvider trace.TracerProvider
	propagator     propagation.TextMapPropagator
	kind           trace.SpanKind
	tracerName     string
	spanName       string
}

// WithPropagator with tracer propagator.
func WithPropagator(propagator propagation.TextMapPropagator) Option {
	return func(opts *options) {
		opts.propagator = propagator
	}
}

// WithTracerProvider with tracer provider.
// By default, it uses the global provider that is set by otel.SetTracerProvider(provider).
func WithTracerProvider(provider trace.TracerProvider) Option {
	return func(opts *options) {
		opts.tracerProvider = provider
	}
}

// WithTracerName with tracer name
func WithTracerName(tracerName string) Option {
	return func(opts *options) {
		opts.tracerName = tracerName
	}
}

// WithGlobalTracerProvider set the registered global trace provider
func WithGlobalTracerProvider() Option {
	return func(opts *options) {
		opts.tracerProvider = otel.GetTracerProvider()
	}
}

// WithGlobalPropagator set the global TextMapPropagator
func WithGlobalPropagator() Option {
	return func(opts *options) {
		opts.propagator = otel.GetTextMapPropagator()
	}
}
