package userbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
)

type UpdateUserStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error)
	Update(ctx context.Context, updatedUser *usermodel.UpdateUser) error
}

type updateUserBiz struct {
	updateUserStore UpdateUserStore
}

func NewUpdateUserBiz(updateUserStore UpdateUserStore) *updateUserBiz {
	return &updateUserBiz{updateUserStore: updateUserStore}
}

func (biz *updateUserBiz) Update(ctx context.Context, filter map[string]interface{}, data *usermodel.UpdateUser) error {

	if err := data.Process(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	existedUser, err := biz.updateUserStore.Find(ctx, filter)
	if err != nil {
		return err
	}

	if existedUser == nil {
		return common.ErrEntityNotFound(data.EntityName(), nil)
	}
	data.Id = existedUser.Id
	err = biz.updateUserStore.Update(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
