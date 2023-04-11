package authbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/firebase"
	"github.com/dinhlockt02/cs_video_call_app_server/components/tokenprovider"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LoginWithFirebaseUserStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error)
}

type LoginWithFirebaseAuthStore interface {
	CreateFirebaseUser(ctx context.Context, data *authmodel.RegisterFirebaseUser) (*primitive.ObjectID, error)
}

type LoginWithFirebaseDeviceStore interface {
	Create(ctx context.Context, data *devicemodel.Device) (*primitive.ObjectID, error)
}

type loginWithFirebaseBiz struct {
	tokenProvider tokenprovider.TokenProvider
	userStore     LoginWithFirebaseUserStore
	deviceStore   LoginWithFirebaseDeviceStore
	fbs           firebase.FirebaseApp
	authStore     LoginWithFirebaseAuthStore
}

func NewLoginWithFirebaseBiz(
	tokenProvider tokenprovider.TokenProvider,
	userStore LoginWithFirebaseUserStore,
	deviceStore LoginWithFirebaseDeviceStore,
	authStore LoginWithFirebaseAuthStore,
	fbs firebase.FirebaseApp,
) *loginWithFirebaseBiz {
	return &loginWithFirebaseBiz{
		tokenProvider: tokenProvider,
		userStore:     userStore,
		deviceStore:   deviceStore,
		fbs:           fbs,
		authStore:     authStore,
	}
}

func (biz *loginWithFirebaseBiz) LoginWithFirebase(ctx context.Context, idToken string, device *devicemodel.Device) (*authmodel.AuthToken, error) {

	if err := device.Process(); err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	uid, err := biz.fbs.VerifyToken(ctx, idToken)
	if err != nil {
		return nil, err
	}

	email, err := biz.fbs.ExtractEmailFromUID(ctx, *uid)

	if err != nil {
		return nil, err
	}
	var userId string
	existedUser, err := biz.userStore.Find(ctx, map[string]interface{}{
		"email": *email,
	})
	if err != nil {
		return nil, err
	}
	if existedUser == nil {

		createdUser := &authmodel.RegisterFirebaseUser{
			Email: *email,
		}

		err = createdUser.Process()
		if err != nil {
			return nil, common.ErrInternal(err)
		}

		t_id, err := biz.authStore.CreateFirebaseUser(ctx, createdUser)
		if err != nil {
			return nil, err
		}

		userId = t_id.Hex()
	} else {
		userId = existedUser.Id.Hex()
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
		&tokenprovider.TokenPayload{UserId: userId},
		common.AccessTokenExpiry,
	)

	return &authmodel.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
