package jwt

import (
	"github.com/dinhlockt02/cs_video_call_app_server/components/tokenprovider"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
)

type TokenProvider struct {
	secret string
}

type AppClaims struct {
	jwt.RegisteredClaims
	Payload *tokenprovider.TokenPayload `json:"payload"`
}

func (j *TokenProvider) Generate(payload *tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	now := time.Now()
	expire := now.Add(time.Second * time.Duration(expiry))
	claims := AppClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Payload: payload,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return nil, errors.Wrap(err, "can not sign token")
	}
	return &tokenprovider.Token{
		Token:     tokenString,
		CreatedAt: &now,
		ExpiredAt: &expire,
	}, nil
}

func (j *TokenProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	tk, err := jwt.ParseWithClaims(
		token,
		&AppClaims{Payload: &tokenprovider.TokenPayload{}},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.secret), nil
		})
	if err != nil {
		return nil, errors.Wrap(err, "can not parse token")
	}

	if claims, ok := tk.Claims.(*AppClaims); ok && tk.Valid {
		return claims.Payload, nil
	}
	return nil, errors.New("invalid token")
}

func NewJwtTokenProvider(secret string) *TokenProvider {
	return &TokenProvider{secret: secret}
}
