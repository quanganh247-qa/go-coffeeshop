package domain

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	List(ctx context.Context) ([]*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) error
}

type RefreshTokenRepository interface {
	Create(ctx context.Context, token *RefreshToken) error
	GetByToken(ctx context.Context, token string) (*RefreshToken, error)
	Revoke(ctx context.Context, token string) error
	RevokeByUserID(ctx context.Context, userID uuid.UUID) error
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

type TokenGenerator interface {
	GenerateAccessToken(user *User) (string, error)
	GenerateRefreshToken(user *User) (string, error)
	VerifyToken(token string) (bool, error)
}
