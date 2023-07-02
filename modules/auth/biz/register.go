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
	"time"
)

type RegisterAuthStore interface {
	CreateEmailAndPasswordUser(ctx context.Context, data *authmodel.RegisterUser) (*authmodel.User, error)
	Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error)
}

type RegisterDeviceStore interface {
	Create(ctx context.Context, data *devicemodel.Device) error
}

type registerBiz struct {
	tokenProvider  tokenprovider.TokenProvider
	passwordHasher hasher.Hasher
	deviceStore    RegisterDeviceStore
	authStore      RegisterAuthStore
}

func NewRegisterBiz(
	tokenProvider tokenprovider.TokenProvider,
	passwordHasher hasher.Hasher,
	deviceStore RegisterDeviceStore,
	authStore RegisterAuthStore,
) *registerBiz {
	return &registerBiz{
		tokenProvider:  tokenProvider,
		passwordHasher: passwordHasher,
		deviceStore:    deviceStore,
		authStore:      authStore,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *authmodel.RegisterUser, device *devicemodel.Device) (*authmodel.AuthToken, error) {

	if err := data.Process(); err != nil {
		return nil, common.ErrInvalidRequest(errors.Wrap(err, "invalid register data"))
	}

	if err := device.Process(); err != nil {
		return nil, common.ErrInvalidRequest(errors.Wrap(err, "invalid device data"))
	}

	existedUser, err := biz.authStore.Find(ctx, authstore.GetEmailFilter(data.Email))
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not find user"))
	}
	if existedUser != nil {
		return nil, common.ErrInvalidRequest(errors.New(authmodel.UserExists))
	}

	hashedPassword, err := biz.passwordHasher.Hash(data.Password)
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not compare password"))
	}

	data.Password = hashedPassword

	user, err := biz.authStore.CreateEmailAndPasswordUser(ctx, data)
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not create user"))
	}

	device.UserId = user.Id
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
		return nil, common.ErrInternal(errors.Wrap(err, "can not create access token"))
	}

	return &authmodel.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
