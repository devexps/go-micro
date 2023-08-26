package engine

import (
	"context"
	"fmt"
	"strings"

	"github.com/devexps/go-micro/v2/errors"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/transport"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
)

// AuthFromMD .
func AuthFromMD(ctx context.Context, ctxType ContextType) (string, error) {
	val := extractTokenFromContext(ctx, ctxType)
	if val == "" {
		return "", errors.Unauthorized(Reason, "Request unauthenticated")
	}
	splits := strings.SplitN(val, " ", 2)
	if len(splits) < 2 {
		return "", errors.Unauthorized(Reason, "Bad authorization string")
	}
	if !strings.EqualFold(splits[0], BearerWord) {
		return "", errors.Unauthorized(Reason, "Request unauthenticated")
	}
	return splits[1], nil
}

// MDWithAuth .
func MDWithAuth(ctx context.Context, tokenStr string, ctxType ContextType) context.Context {
	switch ctxType {
	case ContextTypeGrpc:
		return injectTokenToGrpcContext(ctx, tokenStr)
	case ContextTypeMicro:
		return injectTokenToMicroContext(ctx, tokenStr)
	default:
		return injectTokenToGrpcContext(ctx, tokenStr)
	}
}

func extractTokenFromContext(ctx context.Context, ctxType ContextType) string {
	switch ctxType {
	case ContextTypeGrpc:
		return extractTokenFromGrpcContext(ctx)
	case ContextTypeMicro:
		return extractTokenFromMicroContext(ctx)
	default:
		return extractTokenFromGrpcContext(ctx)
	}
}

func extractTokenFromGrpcContext(ctx context.Context) string {
	return metautils.ExtractIncoming(ctx).Get(AuthorizationKey)
}

func extractTokenFromMicroContext(ctx context.Context) string {
	if header, ok := transport.FromServerContext(ctx); ok {
		return header.RequestHeader().Get(AuthorizationKey)
	}
	return ""
}

func injectTokenToGrpcContext(ctx context.Context, tokenStr string) context.Context {
	metautils.ExtractOutgoing(ctx).Set(AuthorizationKey, formatToken(tokenStr))
	return ctx
}

func injectTokenToMicroContext(ctx context.Context, tokenStr string) context.Context {
	if header, ok := transport.FromClientContext(ctx); ok {
		header.RequestHeader().Set(AuthorizationKey, formatToken(tokenStr))
	} else {
		log.Error("authn token injection failure in go-micro context")
	}
	return ctx
}

func formatToken(tokenStr string) string {
	return fmt.Sprintf(BearerFormat, tokenStr)
}
