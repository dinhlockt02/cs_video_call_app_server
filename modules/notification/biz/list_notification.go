package notibiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	notistore "github.com/dinhlockt02/cs_video_call_app_server/components/notification/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type ListNotificationBiz struct {
	repo notirepo.INotificationRepository
}

func NewListNotificationBiz(repo notirepo.INotificationRepository) *ListNotificationBiz {
	return &ListNotificationBiz{repo: repo}
}

func (biz *ListNotificationBiz) List(ctx context.Context, requesterId string) ([]notimodel.Notification, error) {
	log.Debug().Str("requesterId", requesterId).Msg("list all notifications")
	filter := notistore.GetOwnerFilter(requesterId)
	notis, err := biz.repo.List(ctx, filter)
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not list notifications"))
	}
	return notis, nil
}
