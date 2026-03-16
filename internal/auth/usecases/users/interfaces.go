package usecases

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserRegistrationRequest struct {
	Email    string
	Password string
}

type UserLoginRequest struct {
	Email    string
	Password string
}

type UserLoginResponse struct {
	AccessToken  string
	RefreshToken string
	User         UserDTO
}

type UserDTO struct {
	ID            uuid.UUID
	Email         string
	IsActive      bool
	EmailVerified bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UseCase interface {
	List(ctx context.Context) ([]*UserDTO, error)
	GetByID(ctx context.Context, id uuid.UUID) (*UserDTO, error)
	Login(ctx context.Context, req UserLoginRequest) (*UserLoginResponse, error)
	Register(ctx context.Context, req UserRegistrationRequest) error
	VerifyToken(ctx context.Context, token string) (bool, error)
	RefreshToken(ctx context.Context, token string) (string, error)
}
