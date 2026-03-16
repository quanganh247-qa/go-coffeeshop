package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thangchung/go-coffeeshop/internal/auth/domain"
)

var (
	ErrInvalidToken     = errors.New("invalid token")
	ErrExpiredToken     = errors.New("token has expired")
	ErrInvalidSignature = errors.New("invalid signature")
)

type jwtTokenGenerator struct {
	secretKey string
	issuer    string
	audience  string
}

func NewJWTTokenGenerator(secretKey, issuer, audience string) domain.TokenGenerator {
	return &jwtTokenGenerator{
		secretKey: secretKey,
		issuer:    issuer,
		audience:  audience,
	}
}

func (j *jwtTokenGenerator) GenerateAccessToken(user *domain.User) (string, error) {
	claims := domain.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			Issuer:    j.issuer,
			Audience:  jwt.ClaimStrings{j.audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		Email: user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtTokenGenerator) GenerateRefreshToken(user *domain.User) (string, error) {
	claims := domain.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			Issuer:    j.issuer,
			Audience:  jwt.ClaimStrings{j.audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtTokenGenerator) VerifyToken(tokenString string) (bool, error) {
	claims := &domain.CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSignature
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return false, ErrExpiredToken
		}
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return false, ErrInvalidSignature
		}
		return false, ErrInvalidToken
	}

	if !token.Valid {
		return false, ErrInvalidToken
	}

	return true, nil
}

func (j *jwtTokenGenerator) VerifyAndParseClaims(tokenString string) (*domain.CustomClaims, error) {
	claims := &domain.CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token expired")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
