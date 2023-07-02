package authbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	"github.com/pkg/errors"
)

type IsEmailVerifiedAuthStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error)
}

type IsEmailVerifiedBiz struct {
	authStore LoginAuthStore
}

func NewIsEmailVerifiedBiz(
	authStore IsEmailVerifiedAuthStore,
) *IsEmailVerifiedBiz {
	return &IsEmailVerifiedBiz{
		authStore: authStore,
	}
}

func (biz *IsEmailVerifiedBiz) IsEmailVerified(ctx context.Context, filter map[string]interface{}) (bool, error) {
	existedUser, err := biz.authStore.Find(ctx, filter)
	if err != nil {
		return false, err
	}
	if existedUser == nil {
		return false, common.ErrEntityNotFound(common.UserEntity, errors.New("user not found"))
	}

	return existedUser.EmailVerified, nil
}
