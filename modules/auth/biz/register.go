package authbiz

import (
	"context"
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/hasher"
	"github.com/dinhlockt02/cs_video_call_app_server/components/tokenprovider"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RegisterUserStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error)
}

type RegisterAuthStore interface {
	CreateEmailAndPasswordUser(ctx context.Context, data *authmodel.RegisterUser) (*primitive.ObjectID, error)
}

type RegisterDeviceStore interface {
	Create(ctx context.Context, data *devicemodel.Device) (*primitive.ObjectID, error)
}

type registerBiz struct {
	tokenProvider  tokenprovider.TokenProvider
	userStore      RegisterUserStore
	passwordHasher hasher.Hasher
	deviceStore    RegisterDeviceStore
	authStore      RegisterAuthStore
}

func NewRegisterBiz(
	tokenProvider tokenprovider.TokenProvider,
	userStore RegisterUserStore,
	passwordHasher hasher.Hasher,
	deviceStore RegisterDeviceStore,
	authStore RegisterAuthStore,
) *registerBiz {
	return &registerBiz{
		tokenProvider:  tokenProvider,
		userStore:      userStore,
		passwordHasher: passwordHasher,
		deviceStore:    deviceStore,
		authStore:      authStore,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *authmodel.RegisterUser, device *devicemodel.Device) (*authmodel.AuthToken, error) {

	if err := data.Process(); err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	if err := device.Process(); err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	existedUser, err := biz.userStore.Find(ctx, map[string]interface{}{
		"email": data.Email,
	})
	if err != nil {
		return nil, err
	}
	if existedUser != nil {
		return nil, common.ErrInvalidRequest(errors.New("user existed"))
	}

	hashedPassword, err := biz.passwordHasher.Hash(data.Password)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	data.Password = hashedPassword

	id, err := biz.authStore.CreateEmailAndPasswordUser(ctx, data)
	if err != nil {
		return nil, err
	}

	deviceId, err := biz.deviceStore.Create(ctx, device)

	now := time.Now()
	refreshToken := &tokenprovider.Token{Token: deviceId.Hex(), CreatedAt: &now, ExpiredAt: nil}
	if err != nil {
		return nil, err
	}

	accessToken, err := biz.tokenProvider.Generate(
		&tokenprovider.TokenPayload{UserId: id.Hex()},
		common.AccessTokenExpiry,
	)

	return &authmodel.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
