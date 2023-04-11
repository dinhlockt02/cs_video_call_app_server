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

type LoginUserStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error)
}

type LoginDeviceStore interface {
	Create(ctx context.Context, data *devicemodel.Device) (*primitive.ObjectID, error)
}

type loginBiz struct {
	tokenProvider  tokenprovider.TokenProvider
	userStore      LoginUserStore
	deviceStore    LoginDeviceStore
	passwordHasher hasher.Hasher
}

func NewLoginBiz(
	tokenProvider tokenprovider.TokenProvider,
	userStore LoginUserStore,
	deviceStore LoginDeviceStore,
	passwordHasher hasher.Hasher,
) *loginBiz {
	return &loginBiz{
		tokenProvider:  tokenProvider,
		userStore:      userStore,
		passwordHasher: passwordHasher,
		deviceStore:    deviceStore,
	}
}

func (biz *loginBiz) Login(ctx context.Context, data *authmodel.RegisterUser, device *devicemodel.Device) (*authmodel.AuthToken, error) {

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
	if existedUser == nil {
		return nil, common.ErrInvalidRequest(errors.New("user not exists"))
	}

	isMatch, err := biz.passwordHasher.Compare(data.Password, existedUser.Password)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if !isMatch {
		return nil, common.ErrInvalidRequest(errors.New("invalid email or password"))
	}

	deviceId, err := biz.deviceStore.Create(ctx, device)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	refreshToken := &tokenprovider.Token{Token: deviceId.Hex(), CreatedAt: &now, ExpiredAt: nil}
	if err != nil {
		return nil, err
	}

	accessToken, err := biz.tokenProvider.Generate(
		&tokenprovider.TokenPayload{UserId: existedUser.Id.Hex()},
		common.AccessTokenExpiry,
	)

	return &authmodel.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
