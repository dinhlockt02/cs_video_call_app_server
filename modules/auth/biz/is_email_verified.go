package authbiz

import (
	"context"
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
)

type IsEmailVerifiedAuthStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error)
}

type isEmailVerifiedBiz struct {
	authStore LoginAuthStore
}

func NewIsEmailVerifiedBiz(
	authStore IsEmailVerifiedAuthStore,
) *isEmailVerifiedBiz {
	return &isEmailVerifiedBiz{
		authStore: authStore,
	}
}

func (biz *isEmailVerifiedBiz) IsEmailVerified(ctx context.Context, filter map[string]interface{}) (bool, error) {

	existedUser, err := biz.authStore.Find(ctx, filter)
	if err != nil {
		return false, err
	}
	if existedUser == nil {
		return false, common.ErrInvalidRequest(errors.New("user not exists"))
	}

	return existedUser.EmailVerified, nil
}
