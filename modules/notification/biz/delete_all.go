package notibiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	notistore "github.com/dinhlockt02/cs_video_call_app_server/components/notification/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type DeleteAllNotificationBiz struct {
	repo notirepo.INotificationRepository
}

func NewDeleteAllNotificationBiz(repo notirepo.INotificationRepository) *DeleteAllNotificationBiz {
	return &DeleteAllNotificationBiz{repo: repo}
}

func (biz *DeleteAllNotificationBiz) DeleteAll(ctx context.Context, requesterId string) error {
	log.Debug().Str("requesterId", requesterId).Msg("delete all notifications")
	filter := notistore.GetOwnerFilter(requesterId)
	err := biz.repo.Delete(ctx, filter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not delete all notifications"))
	}
	return nil
}
