package domain

import (
	"time"

	"github.com/google/uuid"
	shared "github.com/thangchung/go-coffeeshop/internal/pkg/shared_kernel"
)

type User struct {
	shared.AggregateRoot
	ID            uuid.UUID
	Email         string
	PasswordHash  string
	IsActive      bool
	EmailVerified bool
	LastLoginAt   *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewUser(email, passwordHash string) *User {
	return &User{
		ID:            uuid.New(),
		Email:         email,
		PasswordHash:  passwordHash,
		IsActive:      true,
		EmailVerified: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

type RefreshToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	ExpiresAt time.Time
	Revoked   bool
	CreatedAt time.Time
}
