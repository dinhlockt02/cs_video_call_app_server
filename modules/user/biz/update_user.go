package userbiz

import (
	"context"
	"encoding/json"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/pubsub"
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
	pubsub          pubsub.PubSub
}

func NewUpdateUserBiz(updateUserStore UpdateUserStore, pubsub pubsub.PubSub) *UpdateUserBiz {
	return &UpdateUserBiz{updateUserStore: updateUserStore, pubsub: pubsub}
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

	biz.publishEvent(ctx, existedUser, data)
	return nil
}

func (biz *UpdateUserBiz) publishEvent(ctx context.Context, existedUser *usermodel.User, data *usermodel.UpdateUser) {
	marshaledUser := &common.User{
		Id: *existedUser.Id,
	}
	if data.Name != nil {
		marshaledUser.Name = *data.Name
	} else {
		marshaledUser.Name = existedUser.Name
	}
	if data.Avatar != nil {
		marshaledUser.Avatar = *data.Avatar
	} else {
		marshaledUser.Avatar = existedUser.Avatar
	}
	marshaledData, err := json.Marshal(marshaledUser)
	if err != nil {
		log.Error().Err(err).Msg("can not marshaled user")
	}

	err = biz.pubsub.Publish(ctx, common.TopicUserUpdateProfile, string(marshaledData))
	if err != nil {
		log.Error().Err(err).Msg("can not publish event")
	}
}
