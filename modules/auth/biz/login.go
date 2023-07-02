package authbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/hasher"
	"github.com/dinhlockt02/cs_video_call_app_server/components/tokenprovider"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	authstore "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/store"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

type LoginAuthStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error)
}

type LoginDeviceStore interface {
	Create(ctx context.Context, data *devicemodel.Device) error
}

type LoginBiz struct {
	tokenProvider  tokenprovider.TokenProvider
	authStore      LoginAuthStore
	deviceStore    LoginDeviceStore
	passwordHasher hasher.Hasher
}

func NewLoginBiz(
	tokenProvider tokenprovider.TokenProvider,
	authStore LoginAuthStore,
	deviceStore LoginDeviceStore,
	passwordHasher hasher.Hasher,
) *LoginBiz {
	return &LoginBiz{
		tokenProvider:  tokenProvider,
		authStore:      authStore,
		passwordHasher: passwordHasher,
		deviceStore:    deviceStore,
	}
}

func (biz *LoginBiz) Login(ctx context.Context, data *authmodel.LoginUser, device *devicemodel.Device) (*authmodel.AuthToken, error) {
	log.Debug().Msg("login usecase executed")
	if err := data.Process(); err != nil {
		return nil, common.ErrInvalidRequest(errors.Wrap(err, "invalid login data"))
	}

	if err := device.Process(); err != nil {
		return nil, common.ErrInvalidRequest(errors.Wrap(err, "invalid device"))
	}

	existedUser, err := biz.authStore.Find(ctx, authstore.GetEmailFilter(data.Email))
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not find a exists user"))
	}
	if existedUser == nil {
		return nil, common.ErrInvalidRequest(errors.New(authmodel.InvalidEmailOrPassword))
	}

	if strings.TrimSpace(existedUser.Password) == "" {
		return nil, common.ErrInvalidRequest(errors.New(authmodel.InvalidEmailOrPassword))
	}

	isMatch, err := biz.passwordHasher.Compare(data.Password, existedUser.Password)
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not compare password"))
	}

	if !isMatch {
		return nil, common.ErrInvalidRequest(errors.New(authmodel.InvalidEmailOrPassword))
	}

	device.UserId = existedUser.Id
	err = biz.deviceStore.Create(ctx, device)
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not create device"))
	}

	now := time.Now()
	refreshToken := &tokenprovider.Token{Token: *device.Id, CreatedAt: &now, ExpiredAt: nil}

	accessToken, err := biz.tokenProvider.Generate(
		&tokenprovider.TokenPayload{Id: *device.Id},
		common.AccessTokenExpiry,
	)

	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not create an access token"))
	}

	return &authmodel.AuthToken{
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		EmailVerified:  existedUser.EmailVerified,
		ProfileUpdated: existedUser.ProfileUpdated,
	}, err
}
