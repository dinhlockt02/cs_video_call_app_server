package jwt

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/tokenprovider"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type jwtTokenProvider struct {
	secret string
}

type AppClaims struct {
	jwt.RegisteredClaims
	payload *tokenprovider.TokenPayload
}

func (j *jwtTokenProvider) Generate(payload *tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	now := time.Now()
	expire := now.Add(time.Second * time.Duration(expiry))
	claims := AppClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
			NotBefore: nil,
			IssuedAt:  jwt.NewNumericDate(now),
		},
		payload: payload,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return &tokenprovider.Token{
		Token:     tokenString,
		CreatedAt: &now,
		ExpiredAt: &expire,
	}, nil
}

func (j *jwtTokenProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	//TODO implement me
	panic("implement me")
}

func NewJwtTokenProvider(secret string) *jwtTokenProvider {
	return &jwtTokenProvider{secret: secret}
}
