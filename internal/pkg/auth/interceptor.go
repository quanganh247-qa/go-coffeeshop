package auth

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Interceptor interface {
	Unary() grpc.UnaryServerInterceptor
}

type interceptor struct {
	tokenVerifier TokenVerifier
	publicMethods []string
}

var _ Interceptor = (*interceptor)(nil)

func NewInterceptor(tokenVerifier TokenVerifier, publicMethods []string) Interceptor {
	return &interceptor{
		tokenVerifier: tokenVerifier,
		publicMethods:  publicMethods,
	}
}

func (i *interceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// 1. Check if method is public
		for _, m := range i.publicMethods {
			if strings.HasPrefix(info.FullMethod, m) {
				return handler(ctx, req)
			}
		}

		// 2. Get token
		token, err := i.getTokenFromMetadata(ctx)
		if err != nil {
			return nil, err
		}

		// 3. Verify token
		claims, err := i.tokenVerifier.VerifyAndParseClaims(token)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid or expired token")
		}

		// 4. Inject into context
		newCtx := context.WithValue(ctx, UserIDKey, claims.UserID)
		newCtx = context.WithValue(newCtx, EmailKey, claims.Email)

		return handler(newCtx, req)
	}
}

func (i *interceptor) getTokenFromMetadata(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "missing metadata")
	}

	authHeaders := md.Get("authorization")
	if len(authHeaders) == 0 {
		return "", status.Error(codes.Unauthenticated, "authorization header is missing")
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeaders[0], bearerPrefix) {
		return "", status.Error(codes.Unauthenticated, "invalid authorization format")
	}

	return strings.TrimPrefix(authHeaders[0], bearerPrefix), nil
}
