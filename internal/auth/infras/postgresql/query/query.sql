-- name: ListUsers :many
SELECT * FROM auth.users;

-- name: GetUserByEmail :one
SELECT * FROM auth.users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM auth.users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO auth.users (email, password_hash, is_active, email_verified, last_login_at, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, email;

-- name: CreateRefreshToken :one
INSERT INTO auth.refresh_tokens (user_id, token, expires_at, revoked, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;