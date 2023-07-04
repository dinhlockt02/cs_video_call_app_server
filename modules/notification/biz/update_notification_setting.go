package notibiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type UpdateNotificationSettingBiz struct {
	repo notirepo.INotificationRepository
}

func NewUpdateNotificationSettingBiz(repo notirepo.INotificationRepository) *UpdateNotificationSettingBiz {
	return &UpdateNotificationSettingBiz{repo: repo}
}

func (biz *UpdateNotificationSettingBiz) Update(ctx context.Context, requesterId string, updatedData *notimodel.NotificationUser) error {
	log.Debug().Str("requesterId", requesterId).
		Any("updatedData", updatedData).
		Msg("update user notification setting")
	userFilter, err := common.GetIdFilter(requesterId)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "invalid requester id"))
	}
	err = biz.repo.UpdateNotificationUser(ctx, userFilter, updatedData)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not get notification user"))
	}

	return nil
}
