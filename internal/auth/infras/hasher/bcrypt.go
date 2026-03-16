package hasher

import (
	"github.com/thangchung/go-coffeeshop/internal/auth/domain"
	"golang.org/x/crypto/bcrypt"
)

type bcryptHasher struct {
	cost int
}

func NewBcryptHasher(cost int) domain.PasswordHasher {
	if cost == 0 {
		cost = bcrypt.DefaultCost
	}
	return &bcryptHasher{cost: cost}
}

func (h *bcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	return string(bytes), err
}

func (h *bcryptHasher) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
