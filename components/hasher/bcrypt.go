package hasher

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

type BcryptHasher struct {
	cost int
}

func NewBcryptHasher(cost ...int) *BcryptHasher {
	var c int
	if len(cost) == 0 {
		c = bcrypt.DefaultCost
	} else if cost[0] < bcrypt.MinCost || cost[0] > bcrypt.MaxCost {
		log.Fatalf("cost must be between %v and % v\n", bcrypt.MinCost, bcrypt.MaxCost)
	}

	return &BcryptHasher{cost: c}
}

func (b *BcryptHasher) Hash(data string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data), b.cost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (b *BcryptHasher) Compare(data string, hashedData string) (bool, error) {
	if strings.TrimSpace(data) == "" {
		return false, errors.New("invalid hashed data")
	}
	if strings.TrimSpace(hashedData) == "" {
		return false, errors.New("invalid hashed data")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashedData), []byte(data)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, errors.Wrap(err, "can not compare hashed data")
	}
	return true, nil
}
