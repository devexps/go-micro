package jwt

import (
	"bytes"
	"strings"

	"github.com/devexps/go-micro/v2/middleware/authn/engine"

	"github.com/golang-jwt/jwt/v4"
)

// ScopeSet see: https://datatracker.ietf.org/doc/html/rfc6749#section-3.3
type ScopeSet map[string]bool

// Claims contains claims that are included in OIDC standard claims.
// See https://openid.net/specs/openid-connect-core-1_0.html#IDToken
type Claims struct {
	Subject string
	Scopes  ScopeSet
}

// GetSubject returns a subject string
func (c *Claims) GetSubject() string {
	return c.Subject
}

// GetScopes returns the scopes object
func (c *Claims) GetScopes() ScopeSet {
	return c.Scopes
}

func mapClaimsToClaims(rawClaims jwt.MapClaims) (engine.Claims, error) {
	// optional subject
	var subject = ""
	if subjectClaim, ok := rawClaims["sub"]; ok {
		if subject, ok = subjectClaim.(string); !ok {
			return nil, ErrInvalidSubject
		}
	}
	claims := &Claims{
		Subject: subject,
		Scopes:  make(ScopeSet),
	}
	// optional scopes
	if scopeKey, ok := rawClaims["scope"]; ok {
		if scope, ok := scopeKey.(string); ok {
			scopes := strings.Split(scope, " ")
			for _, s := range scopes {
				claims.Scopes[s] = true
			}
		}
	}
	return claims, nil
}

func claimsToJwtClaims(rawClaims engine.Claims) jwt.Claims {
	raw := rawClaims.(*Claims)

	claims := jwt.MapClaims{
		"sub": raw.GetSubject(),
	}

	var buffer bytes.Buffer
	count := len(raw.GetScopes())
	idx := 0
	for scope := range raw.GetScopes() {
		buffer.WriteString(scope)
		if idx != count-1 {
			buffer.WriteString(" ")
		}
		idx++
	}
	str := buffer.String()
	if len(str) > 0 {
		claims["scope"] = buffer.String()
	}
	return claims
}
