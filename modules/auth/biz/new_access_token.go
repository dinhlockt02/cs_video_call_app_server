package authbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/tokenprovider"
	devicestore "github.com/dinhlockt02/cs_video_call_app_server/modules/device/store"
	"github.com/pkg/errors"
)

type AccessTokenBiz struct {
	tokenProvider tokenprovider.TokenProvider
	deviceStore   devicestore.Store
}

func NewAccessTokenBiz(
	tokenprovider tokenprovider.TokenProvider,
	deviceStore devicestore.Store,
) *AccessTokenBiz {
	return &AccessTokenBiz{
		tokenProvider: tokenprovider,
		deviceStore:   deviceStore,
	}
}

func (biz *AccessTokenBiz) New(ctx context.Context, refreshToken string) (*tokenprovider.Token, error) {

	deviceFilter, err := common.GetIdFilter(refreshToken)
	if err != nil {
		return nil, common.ErrInvalidRequest(errors.Wrap(err, "invalid refresh token"))
	}

	devices, err := biz.deviceStore.Get(ctx, deviceFilter)
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not find device"))
	}
	if len(devices) == 0 {
		return nil, common.ErrUnauthorized(errors.New("invalid refresh token"))
	}

	accessToken, err := biz.tokenProvider.Generate(
		&tokenprovider.TokenPayload{Id: refreshToken},
		common.AccessTokenExpiry,
	)

	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not regenerate access token"))
	}

	return accessToken, nil
}
