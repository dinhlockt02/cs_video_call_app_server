package authbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/hasher"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	"github.com/pkg/errors"
)

type UpdatePasswordAuthStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error)
	Update(ctx context.Context, filter map[string]interface{}, passwordUser *authmodel.UpdatePasswordUser) error
}

type UpdatePasswordBiz struct {
	authStore      UpdatePasswordAuthStore
	passwordHasher hasher.Hasher
}

func NewUpdatePasswordBiz(
	authStore UpdatePasswordAuthStore,
	passwordHasher hasher.Hasher,
) *UpdatePasswordBiz {
	return &UpdatePasswordBiz{
		authStore:      authStore,
		passwordHasher: passwordHasher,
	}
}

func (biz *UpdatePasswordBiz) Update(ctx context.Context,
	filter map[string]interface{}, data *authmodel.UpdatePasswordUser) error {
	if err := data.Process(); err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid update data"))
	}

	existedUser, err := biz.authStore.Find(ctx, filter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find user"))
	}
	if existedUser == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.New(authmodel.UserNotFound))
	}

	hashedPassword, err := biz.passwordHasher.Hash(data.Password)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not hash password"))
	}

	data.Password = hashedPassword

	err = biz.authStore.Update(ctx, filter, data)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not update password"))
	}

	return nil
}
