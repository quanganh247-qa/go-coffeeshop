package token

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/thangchung/go-coffeeshop/internal/auth/domain"
)

type jwtTokenGenerator struct {
	secretKey string
}

func NewJWTTokenGenerator(secretKey string) domain.TokenGenerator {
	return &jwtTokenGenerator{secretKey: secretKey}
}

func (j *jwtTokenGenerator) GenerateAccessToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID.String(),
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtTokenGenerator) GenerateRefreshToken(user *domain.User) (string, error) {
	// Simple for now, can be more complex if needed
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtTokenGenerator) VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil || !token.Valid {
		return false, nil
	}

	return true, nil
}
