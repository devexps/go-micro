package mqtt

import "github.com/devexps/go-micro/v2/broker"

///
/// Option
///

type cleanSessionKey struct{}
type authKey struct{}
type clientIdKey struct{}
type autoReconnectKey struct{}
type resumeSubsKey struct{}
type orderMattersKey struct{}
type errorLoggerKey struct{}
type criticalLoggerKey struct{}
type warnLoggerKey struct{}
type debugLoggerKey struct{}
type loggerKey struct{}

type AuthRecord struct {
	Username string
	Password string
}

type LoggerOptions struct {
	Error    bool
	Critical bool
	Warn     bool
	Debug    bool
}

// WithCleanSession enable clean session option
func WithCleanSession(enable bool) broker.Option {
	return broker.WithContextAndValue(cleanSessionKey{}, enable)
}

// WithAuth set username & password options
func WithAuth(username string, password string) broker.Option {
	return broker.WithContextAndValue(authKey{}, &AuthRecord{
		Username: username,
		Password: password,
	})
}

// WithClientId set client id option
func WithClientId(clientId string) broker.Option {
	return broker.WithContextAndValue(clientIdKey{}, clientId)
}

// WithAutoReconnect enable aut reconnect option
func WithAutoReconnect(enable bool) broker.Option {
	return broker.WithContextAndValue(autoReconnectKey{}, enable)
}

// WithResumeSubs .
func WithResumeSubs(enable bool) broker.Option {
	return broker.WithContextAndValue(resumeSubsKey{}, enable)
}

// WithOrderMatters .
func WithOrderMatters(enable bool) broker.Option {
	return broker.WithContextAndValue(orderMattersKey{}, enable)
}

// WithErrorLogger .
func WithErrorLogger() broker.Option {
	return broker.WithContextAndValue(errorLoggerKey{}, true)
}

// WithCriticalLogger .
func WithCriticalLogger() broker.Option {
	return broker.WithContextAndValue(criticalLoggerKey{}, true)
}

// WithWarnLogger .
func WithWarnLogger() broker.Option {
	return broker.WithContextAndValue(warnLoggerKey{}, true)
}

// WithDebugLogger .
func WithDebugLogger() broker.Option {
	return broker.WithContextAndValue(debugLoggerKey{}, true)
}

// WithLogger .
func WithLogger(opt LoggerOptions) broker.Option {
	return broker.WithContextAndValue(loggerKey{}, opt)
}

///
/// SubscribeOption
///

type qosSubscribeKey struct{}

// WithSubscribeQos QOS
func WithSubscribeQos(qos byte) broker.SubscribeOption {
	return broker.WithSubscribeContextAndValue(qosSubscribeKey{}, qos)
}

///
/// PublishOption
///

type qosPublishKey struct{}
type retainedPublishKey struct{}

// WithPublishQos QOS
func WithPublishQos(qos byte) broker.PublishOption {
	return broker.WithPublishContextAndValue(qosPublishKey{}, qos)
}

// WithPublishRetained retained
func WithPublishRetained(qos byte) broker.PublishOption {
	return broker.WithPublishContextAndValue(retainedPublishKey{}, qos)
}
