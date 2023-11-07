package jwt

import (
	"context"

	"github.com/devexps/go-micro/v2/middleware/authn/engine"

	jwtSdk "github.com/golang-jwt/jwt/v4"
)

type authenticator struct {
	opts *options
}

var _ engine.Authenticator = (*authenticator)(nil)

// NewAuthenticator creates a new authenticator with custom options
func NewAuthenticator(opts ...Option) (engine.Authenticator, error) {
	auth := &authenticator{
		opts: &options{},
	}
	for _, o := range opts {
		o(auth.opts)
	}
	if auth.opts.signingMethod == nil {
		auth.opts.signingMethod = jwtSdk.SigningMethodHS256
	}
	return auth, nil
}

// Authenticate .
func (a *authenticator) Authenticate(ctx context.Context, ctxType engine.ContextType) (engine.Claims, error) {
	tokenString, err := engine.AuthFromMD(ctx, ctxType)
	if err != nil {
		return nil, ErrMissingBearerToken
	}
	return a.authenticate(tokenString)
}

// CreateIdentityWithContext .
func (a *authenticator) CreateIdentityWithContext(ctx context.Context, ctxType engine.ContextType, claims engine.Claims) (context.Context, error) {
	strToken, err := a.CreateIdentity(claims)
	if err != nil {
		return ctx, err
	}
	ctx = engine.MDWithAuth(ctx, strToken, ctxType)
	return ctx, nil
}

// CreateIdentity .
func (a *authenticator) CreateIdentity(claims engine.Claims) (string, error) {
	jwtToken := jwtSdk.NewWithClaims(a.opts.signingMethod, claimsToJwtClaims(claims))

	strToken, err := a.generateToken(jwtToken)
	if err != nil {
		return "", err
	}
	return strToken, nil
}

func (a *authenticator) authenticate(token string) (engine.Claims, error) {
	jwtToken, err := a.parseToken(token)
	if err != nil {
		ve, ok := err.(*jwtSdk.ValidationError)
		if !ok {
			return nil, ErrUnauthenticated
		}
		if ve.Errors&jwtSdk.ValidationErrorMalformed != 0 {
			return nil, ErrInvalidToken
		}
		if ve.Errors&(jwtSdk.ValidationErrorExpired|jwtSdk.ValidationErrorNotValidYet) != 0 {
			return nil, ErrTokenExpired
		}
		return nil, ErrInvalidToken
	}
	if !jwtToken.Valid {
		return nil, ErrInvalidToken
	}
	if jwtToken.Method != a.opts.signingMethod {
		return nil, ErrUnsupportedSigningMethod
	}
	if jwtToken.Claims == nil {
		return nil, ErrInvalidClaims
	}
	mapClaims, ok := jwtToken.Claims.(jwtSdk.MapClaims)
	if !ok {
		return nil, ErrInvalidClaims
	}
	claims, err := mapClaimsToClaims(mapClaims)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func (a *authenticator) parseToken(token string) (*jwtSdk.Token, error) {
	if a.opts.keyFunc == nil {
		return nil, ErrMissingKeyFunc
	}
	return jwtSdk.Parse(token, a.opts.keyFunc)
}

func (a *authenticator) generateToken(jwtToken *jwtSdk.Token) (string, error) {
	if a.opts.keyFunc == nil {
		return "", ErrMissingKeyFunc
	}
	key, err := a.opts.keyFunc(jwtToken)
	if err != nil {
		return "", ErrGetKeyFailed
	}
	strToken, err := jwtToken.SignedString(key)
	if err != nil {
		return "", ErrSignTokenFailed
	}
	return strToken, nil
}
