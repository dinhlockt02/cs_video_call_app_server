package notibiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	notistore "github.com/dinhlockt02/cs_video_call_app_server/components/notification/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type DeleteByIdNotificationBiz struct {
	repo notirepo.INotificationRepository
}

func NewDeleteByIdNotificationBiz(repo notirepo.INotificationRepository) *DeleteByIdNotificationBiz {
	return &DeleteByIdNotificationBiz{repo: repo}
}

func (biz *DeleteByIdNotificationBiz) DeleteById(ctx context.Context, requesterId, id string) error {
	log.Debug().Str("requesterId", requesterId).Str("id", id).Msg("delete notification by id")
	filter, err := common.GetIdFilter(id)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid notification id"))
	}
	filter = common.GetAndFilter(filter, notistore.GetOwnerFilter(requesterId))
	err = biz.repo.Delete(ctx, filter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not delete notification by id"))
	}
	return nil
}
