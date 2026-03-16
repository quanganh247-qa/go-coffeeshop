package repo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/thangchung/go-coffeeshop/internal/auth/domain"
	"github.com/thangchung/go-coffeeshop/internal/auth/infras/postgresql"
)

var _ domain.RefreshTokenRepository = (*refreshTokenRepository)(nil)

type refreshTokenRepository struct {
	db *sql.DB
	q  *postgresql.Queries
}

func NewRefreshTokenRepository(db *sql.DB) domain.RefreshTokenRepository {
	return &refreshTokenRepository{
		db: db,
		q:  postgresql.New(db),
	}
}

func (r *refreshTokenRepository) Create(ctx context.Context, token *domain.RefreshToken) error {
	_, err := r.q.CreateRefreshToken(ctx, postgresql.CreateRefreshTokenParams{
		UserID:    uuid.NullUUID{UUID: token.UserID, Valid: true},
		Token:     token.Token,
		ExpiresAt: token.ExpiresAt,
		Revoked:   sql.NullBool{Bool: token.Revoked, Valid: true},
		CreatedAt: sql.NullTime{Time: token.CreatedAt, Valid: true},
	})
	return err
}

func (r *refreshTokenRepository) GetByToken(ctx context.Context, token string) (*domain.RefreshToken, error) {
	return nil, nil
}

func (r *refreshTokenRepository) Revoke(ctx context.Context, token string) error {
	return nil
}

func (r *refreshTokenRepository) RevokeByUserID(ctx context.Context, userID uuid.UUID) error {
	return nil
}
