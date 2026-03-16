package auth

import (
	"context"

	jwt "github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const (
	UserIDKey ctxKey = "user_id"
	EmailKey  ctxKey = "user_email"
)

// CustomClaims represents the JWT claims
type CustomClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"sub"`
	Email  string `json:"email,omitempty"`
}

// TokenVerifier defines the interface for verifying tokens
type TokenVerifier interface {
	VerifyAndParseClaims(token string) (*CustomClaims, error)
}

// UserIDFromContext retrieves the user ID from the context
func UserIDFromContext(ctx context.Context) string {
	if id, ok := ctx.Value(UserIDKey).(string); ok {
		return id
	}
	return ""
}

// EmailFromContext retrieves the email from the context
func EmailFromContext(ctx context.Context) string {
	if email, ok := ctx.Value(EmailKey).(string); ok {
		return email
	}
	return ""
}
