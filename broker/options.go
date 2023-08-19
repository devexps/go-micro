package broker

import (
	"context"
	"crypto/tls"
	"github.com/devexps/go-micro/v2/encoding"
	"github.com/devexps/go-micro/v2/tracing"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

var (
	DefaultCodec encoding.Codec = nil
)

///////////////////////////////////////////////////////////////////////////////

type Options struct {
	Addrs        []string
	Codec        encoding.Codec
	ErrorHandler Handler
	Secure       bool
	TLSConfig    *tls.Config
	Context      context.Context
	Tracings     []tracing.Option
}

type Option func(*Options)

// NewOptions .
func NewOptions(opts ...Option) Options {
	opt := Options{
		Addrs:        []string{},
		Codec:        DefaultCodec,
		ErrorHandler: nil,
		Secure:       false,
		TLSConfig:    nil,
		Context:      context.Background(),
		Tracings:     []tracing.Option{},
	}
	opt.Apply(opts...)
	return opt
}

// Apply .
func (o *Options) Apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithContext .
func WithContext(ctx context.Context) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = ctx
		}
	}
}

// WithContextAndValue .
func WithContextAndValue(k, v interface{}) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

// WithAddress set broker address
func WithAddress(addressList ...string) Option {
	return func(o *Options) {
		o.Addrs = addressList
	}
}

// WithCodec set codec, support: json, proto.
func WithCodec(name string) Option {
	return func(o *Options) {
		o.Codec = encoding.GetCodec(name)
	}
}

// WithErrorHandler .
func WithErrorHandler(handler Handler) Option {
	return func(o *Options) {
		o.ErrorHandler = handler
	}
}

// WithEnableSecure .
func WithEnableSecure(enable bool) Option {
	return func(o *Options) {
		o.Secure = enable
	}
}

// WithTLSConfig .
func WithTLSConfig(config *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = config
		if o.TLSConfig != nil {
			o.Secure = true
		}
	}
}

// WithTracerProvider .
func WithTracerProvider(provider trace.TracerProvider, tracerName string) Option {
	return func(opt *Options) {
		opt.Tracings = append(opt.Tracings, tracing.WithTracerProvider(provider))
	}
}

// WithPropagator .
func WithPropagator(propagators propagation.TextMapPropagator) Option {
	return func(opt *Options) {
		opt.Tracings = append(opt.Tracings, tracing.WithPropagator(propagators))
	}
}

// WithGlobalTracerProvider .
func WithGlobalTracerProvider() Option {
	return func(opt *Options) {
		opt.Tracings = append(opt.Tracings, tracing.WithGlobalTracerProvider())
	}
}

// WithGlobalPropagator .
func WithGlobalPropagator() Option {
	return func(opt *Options) {
		opt.Tracings = append(opt.Tracings, tracing.WithGlobalPropagator())
	}
}

///////////////////////////////////////////////////////////////////////////////

type PublishOptions struct {
	Context context.Context
}

type PublishOption func(*PublishOptions)

// NewPublishOptions .
func NewPublishOptions(opts ...PublishOption) PublishOptions {
	opt := PublishOptions{
		Context: context.Background(),
	}
	opt.Apply(opts...)
	return opt
}

// Apply .
func (o *PublishOptions) Apply(opts ...PublishOption) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithPublishContext .
func WithPublishContext(ctx context.Context) PublishOption {
	return func(o *PublishOptions) {
		o.Context = ctx
	}
}

// WithPublishContextAndValue .
func WithPublishContextAndValue(k, v interface{}) PublishOption {
	return func(o *PublishOptions) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

///////////////////////////////////////////////////////////////////////////////

type SubscribeOptions struct {
	AutoAck bool
	Queue   string
	Context context.Context
}

type SubscribeOption func(*SubscribeOptions)

// NewSubscribeOptions .
func NewSubscribeOptions(opts ...SubscribeOption) SubscribeOptions {
	opt := SubscribeOptions{
		AutoAck: true,
		Queue:   "",
		Context: context.Background(),
	}
	opt.Apply(opts...)
	return opt
}

// Apply .
func (o *SubscribeOptions) Apply(opts ...SubscribeOption) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithSubscribeContext .
func WithSubscribeContext(ctx context.Context) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Context = ctx
	}
}

// WithSubscribeContextAndValue .
func WithSubscribeContextAndValue(k, v interface{}) SubscribeOption {
	return func(o *SubscribeOptions) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

// WithQueueName .
func WithQueueName(name string) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Queue = name
	}
}

// DisableAutoAck .
func DisableAutoAck() SubscribeOption {
	return func(o *SubscribeOptions) {
		o.AutoAck = false
	}
}
