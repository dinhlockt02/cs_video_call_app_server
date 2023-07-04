package notibiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type GetNotificationSettingBiz struct {
	repo notirepo.INotificationRepository
}

func NewGetNotificationSettingBiz(repo notirepo.INotificationRepository) *GetNotificationSettingBiz {
	return &GetNotificationSettingBiz{repo: repo}
}

func (biz *GetNotificationSettingBiz) Get(ctx context.Context, requesterId string) (bool, error) {
	log.Debug().Str("requesterId", requesterId).Msg("get user notification setting")
	userFilter, err := common.GetIdFilter(requesterId)
	if err != nil {
		return false, common.ErrInternal(errors.Wrap(err, "invalid requester id"))
	}
	user, err := biz.repo.GetNotificationUser(ctx, userFilter)
	if err != nil {
		return false, common.ErrInternal(errors.Wrap(err, "can not get notification user"))
	}
	if user == nil {
		return false, common.ErrInternal(errors.Wrap(err, "requester not found"))
	}

	return user.Notification, nil
}
