package tracing

import (
	"context"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware"
	"github.com/devexps/go-micro/v2/transport"
	"go.opentelemetry.io/otel/attribute"
	semConv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	defaultServerSpanName = "go-micro-server"
	defaultClientSpanName = "go-micro-client"
)

// Server returns a new server middleware for OpenTelemetry.
func Server(opts ...Option) middleware.Middleware {
	tracer := NewTracer(trace.SpanKindServer, defaultServerSpanName, opts...)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				var span trace.Span
				attrs := []attribute.KeyValue{
					semConv.MessagingOperationKey.String(tr.Operation()),
				}
				ctx, span = tracer.Start(ctx, tr.RequestHeader(), attrs...)
				setServerSpan(ctx, span, req)
				defer func() { tracer.End(ctx, span, err) }()
			}
			return handler(ctx, req)
		}
	}
}

// Client returns a new client middleware for OpenTelemetry.
func Client(opts ...Option) middleware.Middleware {
	tracer := NewTracer(trace.SpanKindClient, defaultClientSpanName, opts...)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromClientContext(ctx); ok {
				var span trace.Span
				attrs := []attribute.KeyValue{
					semConv.MessagingOperationKey.String(tr.Operation()),
				}
				ctx, span = tracer.Start(ctx, tr.RequestHeader(), attrs...)
				setClientSpan(ctx, span, req)
				defer func() { tracer.End(ctx, span, err) }()
			}
			return handler(ctx, req)
		}
	}
}

// TraceID returns a traceid valuer.
func TraceID() log.Valuer {
	return func(ctx context.Context) interface{} {
		if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
			return span.TraceID().String()
		}
		return ""
	}
}

// SpanID returns a spanid valuer.
func SpanID() log.Valuer {
	return func(ctx context.Context) interface{} {
		if span := trace.SpanContextFromContext(ctx); span.HasSpanID() {
			return span.SpanID().String()
		}
		return ""
	}
}
