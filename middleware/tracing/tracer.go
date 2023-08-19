package tracing

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const defaultTracerName = "go-micro"

// Tracer is otel span tracer
type Tracer struct {
	tracer trace.Tracer
	opt    *options
}

// NewTracer create tracer instance
func NewTracer(kind trace.SpanKind, spanName string, opts ...Option) *Tracer {
	op := options{
		propagator: propagation.NewCompositeTextMapPropagator(Metadata{}, propagation.Baggage{}, propagation.TraceContext{}),
		kind:       kind,
		tracerName: defaultTracerName,
		spanName:   spanName,
	}
	for _, o := range opts {
		o(&op)
	}
	if op.tracerProvider != nil {
		otel.SetTracerProvider(op.tracerProvider)
	}

	switch kind {
	case trace.SpanKindProducer, trace.SpanKindConsumer:
		return &Tracer{tracer: otel.Tracer(op.tracerName), opt: &op}
	case trace.SpanKindServer, trace.SpanKindClient:
		return &Tracer{tracer: otel.Tracer(op.tracerName), opt: &op}
	default:
		panic(fmt.Sprintf("unsupported span kind: %v", kind))
	}
}

// Start starts tracing span
func (t *Tracer) Start(ctx context.Context, carrier propagation.TextMapCarrier, attrs ...attribute.KeyValue) (context.Context, trace.Span) {
	if t.opt.kind == trace.SpanKindServer || t.opt.kind == trace.SpanKindConsumer {
		ctx = t.opt.propagator.Extract(ctx, carrier)
	}
	opts := []trace.SpanStartOption{
		trace.WithAttributes(attrs...),
		trace.WithSpanKind(t.opt.kind),
	}
	ctx, span := t.tracer.Start(ctx, t.opt.spanName, opts...)
	if t.opt.kind == trace.SpanKindClient || t.opt.kind == trace.SpanKindProducer {
		t.Inject(ctx, carrier)
	}
	return ctx, span
}

// End finish tracing span
func (t *Tracer) End(_ context.Context, span trace.Span, err error, attrs ...attribute.KeyValue) {
	if span == nil {
		return
	}
	if !span.IsRecording() {
		return
	}
	span.SetAttributes(attrs...)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	}
	span.End()
}

// Inject set cross-cutting concerns from the Context into the carrier.
func (t *Tracer) Inject(ctx context.Context, carrier propagation.TextMapCarrier) {
	t.opt.propagator.Inject(ctx, carrier)
}
