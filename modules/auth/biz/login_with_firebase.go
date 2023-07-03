package authbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/firebase"
	"github.com/dinhlockt02/cs_video_call_app_server/components/tokenprovider"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	authstore "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/store"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"time"
)

type LoginWithFirebaseAuthStore interface {
	CreateFirebaseUser(ctx context.Context, data *authmodel.RegisterFirebaseUser) (*authmodel.User, error)
	Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error)
	DeleteUser(ctx context.Context, filter map[string]interface{}) error
}

type LoginWithFirebaseDeviceStore interface {
	Create(ctx context.Context, data *devicemodel.Device) error
}

type LoginWithFirebaseBiz struct {
	tokenProvider tokenprovider.TokenProvider
	deviceStore   LoginWithFirebaseDeviceStore
	fbs           firebase.App
	authStore     LoginWithFirebaseAuthStore
}

func NewLoginWithFirebaseBiz(
	tokenProvider tokenprovider.TokenProvider,
	deviceStore LoginWithFirebaseDeviceStore,
	authStore LoginWithFirebaseAuthStore,
	fbs firebase.App,
) *LoginWithFirebaseBiz {
	return &LoginWithFirebaseBiz{
		tokenProvider: tokenProvider,
		deviceStore:   deviceStore,
		fbs:           fbs,
		authStore:     authStore,
	}
}

func (biz *LoginWithFirebaseBiz) LoginWithFirebase(ctx context.Context,
	idToken string, device *devicemodel.Device) (*authmodel.AuthToken, error) {
	if err := device.Process(); err != nil {
		return nil, common.ErrInvalidRequest(errors.Wrap(err, "invalid device data"))
	}

	uid, err := biz.fbs.VerifyToken(ctx, idToken)
	if err != nil {
		// TODO: handle cases that firebase app is down
		return nil, common.ErrInvalidRequest(errors.Wrap(err, "can not verify token"))
	}

	email, err := biz.fbs.ExtractEmailFromUID(ctx, *uid)

	if err != nil {
		// TODO: handle cases that firebase app is down
		return nil, common.ErrInternal(errors.Wrap(err, "can not extract email from uid"))
	}
	existedUser, err := biz.authStore.Find(ctx, authstore.GetEmailFilter(*email))

	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not get user"))
	}

	if existedUser != nil && !existedUser.EmailVerified {
		log.Warn().Str("email", existedUser.Email).Msg("delete user as they has not verified email")
		err = biz.authStore.DeleteUser(ctx, map[string]interface{}{
			"email": *email,
		})

		if err != nil {
			return nil, common.ErrInternal(errors.Wrap(err, "can not delete user"))
		}
		existedUser = nil
	}

	if existedUser == nil {
		createdUser := &authmodel.RegisterFirebaseUser{
			Email: *email,
		}
		err = createdUser.Process()
		if err != nil {
			return nil, common.ErrInternal(errors.Wrap(err, "can not process created user"))
		}

		existedUser, err = biz.authStore.CreateFirebaseUser(ctx, createdUser)
		if err != nil {
			return nil, common.ErrInternal(errors.Wrap(err, "can not create user"))
		}

	}

	device.UserId = existedUser.Id
	err = biz.deviceStore.Create(ctx, device)
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not create device session"))
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
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		EmailVerified:  existedUser.EmailVerified,
		ProfileUpdated: existedUser.ProfileUpdated,
	}, err
}
