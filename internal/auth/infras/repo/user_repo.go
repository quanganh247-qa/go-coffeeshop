package repo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/thangchung/go-coffeeshop/internal/auth/domain"
	"github.com/thangchung/go-coffeeshop/internal/auth/infras/postgresql"
)

var _ domain.UserRepository = (*userRepository)(nil)

var RepositorySet = wire.NewSet(NewUserRepository, NewRefreshTokenRepository)

type userRepository struct {
	db *sql.DB
	q  *postgresql.Queries
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		db: db,
		q:  postgresql.New(db),
	}
}

func (r *userRepository) List(ctx context.Context) ([]*domain.User, error) {
	users, err := r.q.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.User, 0, len(users))
	for _, u := range users {
		result = append(result, &domain.User{
			ID:            u.ID,
			Email:         u.Email,
			PasswordHash:  u.PasswordHash,
			IsActive:      u.IsActive.Bool,
			EmailVerified: u.EmailVerified.Bool,
			LastLoginAt:   &u.LastLoginAt.Time,
			CreatedAt:     u.CreatedAt,
			UpdatedAt:     u.UpdatedAt,
		})
	}
	return result, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	u, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &domain.User{
		ID:            u.ID,
		Email:         u.Email,
		PasswordHash:  u.PasswordHash,
		IsActive:      u.IsActive.Bool,
		EmailVerified: u.EmailVerified.Bool,
		LastLoginAt:   &u.LastLoginAt.Time,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}, nil
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	u, err := r.q.GetUserByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &domain.User{
		ID:            u.ID,
		Email:         u.Email,
		PasswordHash:  u.PasswordHash,
		IsActive:      u.IsActive.Bool,
		EmailVerified: u.EmailVerified.Bool,
		LastLoginAt:   &u.LastLoginAt.Time,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}, nil
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	newUser, err := r.q.CreateUser(ctx, postgresql.CreateUserParams{
		Email:         user.Email,
		PasswordHash:  user.PasswordHash,
		IsActive:      sql.NullBool{Bool: user.IsActive, Valid: true},
		EmailVerified: sql.NullBool{Bool: user.EmailVerified, Valid: true},
		LastLoginAt:   sql.NullTime{Valid: false},
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}
	user.ID = newUser.ID
	user.Email = newUser.Email
	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	// Need to add UpdateUser to query.sql
	return nil
}
