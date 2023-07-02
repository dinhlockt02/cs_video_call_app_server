package authbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/tokenprovider"
	"github.com/pkg/errors"
)

type accessTokenBiz struct {
	tokenProvider tokenprovider.TokenProvider
}

func NewAccessTokenBiz(
	tokenprovider tokenprovider.TokenProvider,
) *accessTokenBiz {
	return &accessTokenBiz{
		tokenProvider: tokenprovider,
	}
}

func (biz *accessTokenBiz) New(ctx context.Context, refreshToken string) (*tokenprovider.Token, error) {

	accessToken, err := biz.tokenProvider.Generate(
		&tokenprovider.TokenPayload{Id: refreshToken},
		common.AccessTokenExpiry,
	)

	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not regenerate access token"))
	}

	return accessToken, nil
}
