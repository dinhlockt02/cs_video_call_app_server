package hasher

import (
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

type bcryptHasher struct {
	cost int
}

func NewBcryptHasher(cost ...int) *bcryptHasher {
	var c int
	if len(cost) == 0 {
		c = bcrypt.DefaultCost
	} else if cost[0] < bcrypt.MinCost || cost[0] > bcrypt.MaxCost {
		log.Fatalf("cost must be between %v and % v\n", bcrypt.MinCost, bcrypt.MaxCost)
	}

	return &bcryptHasher{cost: c}
}

func (b *bcryptHasher) Hash(data string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data), b.cost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (b *bcryptHasher) Compare(data string, hashedData string) (bool, error) {
	if len(data) < 8 || len(data) > 50 {
		return false, common.ErrInvalidRequest(errors.New("invalid raw data"))
	}
	if strings.TrimSpace(hashedData) == "" {
		return false, common.ErrInvalidRequest(errors.New("invalid hashed data"))
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashedData), []byte(data)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
