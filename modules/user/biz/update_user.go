package userbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type UpdateUserStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error)
	Update(ctx context.Context, filter map[string]interface{}, updatedUser *usermodel.UpdateUser) error
}

type UpdateUserBiz struct {
	updateUserStore UpdateUserStore
}

func NewUpdateUserBiz(updateUserStore UpdateUserStore) *UpdateUserBiz {
	return &UpdateUserBiz{updateUserStore: updateUserStore}
}

func (biz *UpdateUserBiz) Update(ctx context.Context, filter map[string]interface{}, data *usermodel.UpdateUser) error {
	log.Debug().Any("data", data).Any("filter", filter).Msg("find user")
	if err := data.Process(); err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid update user data"))
	}

	existedUser, err := biz.updateUserStore.Find(ctx, filter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find user"))
	}

	if existedUser == nil {
		return common.ErrEntityNotFound(common.UserEntity, usermodel.ErrUserNotFound)
	}
	err = biz.updateUserStore.Update(ctx, filter, data)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not update user"))
	}
	return nil
}
