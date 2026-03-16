package usecases

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/thangchung/go-coffeeshop/internal/auth/domain"
)

type userService struct {
	userRepo         domain.UserRepository
	refreshTokenRepo domain.RefreshTokenRepository
	hasher           domain.PasswordHasher
	tokenGenerator   domain.TokenGenerator
}

func NewUserService(
	userRepo domain.UserRepository,
	refreshTokenRepo domain.RefreshTokenRepository,
	hasher domain.PasswordHasher,
	tokenGenerator domain.TokenGenerator,
) UseCase {
	return &userService{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		hasher:           hasher,
		tokenGenerator:   tokenGenerator,
	}
}

func (s *userService) List(ctx context.Context) ([]*UserDTO, error) {
	users, err := s.userRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*UserDTO, 0, len(users))
	for _, u := range users {
		result = append(result, s.toUserDTO(u))
	}
	return result, nil
}

func (s *userService) GetByID(ctx context.Context, id uuid.UUID) (*UserDTO, error) {
	u, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("user not found")
	}

	return s.toUserDTO(u), nil
}

func (s *userService) Register(ctx context.Context, req UserRegistrationRequest) error {
	existingUser, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := s.hasher.Hash(req.Password)
	if err != nil {
		return err
	}

	user := domain.NewUser(req.Email, hashedPassword)
	_, err = s.userRepo.Create(ctx, user)
	return err
}

func (s *userService) Login(ctx context.Context, req UserLoginRequest) (*UserLoginResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	err = s.hasher.Compare(user.PasswordHash, req.Password)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	accessToken, err := s.tokenGenerator.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenGenerator.GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	return &UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         *s.toUserDTO(user),
	}, nil
}

func (s *userService) VerifyToken(ctx context.Context, tokenString string) (bool, error) {
	return s.tokenGenerator.VerifyToken(tokenString)
}

func (s *userService) RefreshToken(ctx context.Context, token string) (string, error) {
	valid, err := s.tokenGenerator.VerifyToken(token)
	if err != nil || !valid {
		return "", errors.New("invalid refresh token")
	}

	// In a real implementation, you would look up the user by token or claims
	// and maybe check if the token is revoked in the database.
	// For this exercise, we'll keep it simple.

	return "", errors.New("refresh token not fully implemented")
}

func (s *userService) toUserDTO(user *domain.User) *UserDTO {
	return &UserDTO{
		ID:            user.ID,
		Email:         user.Email,
		IsActive:      user.IsActive,
		EmailVerified: user.EmailVerified,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}
