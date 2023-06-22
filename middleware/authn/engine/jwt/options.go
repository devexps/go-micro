package jwt

import "github.com/golang-jwt/jwt/v4"

// Option is jwt option
type Option func(*options)

type options struct {
	signingMethod jwt.SigningMethod
	keyFunc       jwt.Keyfunc
}

// WithSigningMethod set signing method
func WithSigningMethod(alg string) Option {
	return func(o *options) {
		o.signingMethod = jwt.GetSigningMethod(alg)
	}
}

// WithKey set key
func WithKey(key []byte) Option {
	return func(o *options) {
		o.keyFunc = func(token *jwt.Token) (interface{}, error) {
			return key, nil
		}
	}
}
