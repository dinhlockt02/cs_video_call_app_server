package messagebiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	messagerepo "github.com/dinhlockt02/cs_video_call_app_server/modules/message/repository"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type UnpinMessageBiz struct {
	messageRepo messagerepo.Repository
}

func NewUnpinMessageBiz(messageRepo messagerepo.Repository) *UnpinMessageBiz {
	return &UnpinMessageBiz{messageRepo: messageRepo}
}

func (biz *UnpinMessageBiz) Unpin(ctx context.Context, requesterId, messageId string) error {
	log.Debug().Str("requesterId", requesterId).Str("messageId", messageId).Msg("unpin message")
	messageFilter, err := common.GetIdFilter(messageId)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid message id"))
	}

	err = biz.messageRepo.DeleteMessage(ctx, messageFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not delete sender"))
	}
	return nil
}
