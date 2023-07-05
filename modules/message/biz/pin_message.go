package messagebiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	messagemdl "github.com/dinhlockt02/cs_video_call_app_server/modules/message/model"
	messagerepo "github.com/dinhlockt02/cs_video_call_app_server/modules/message/repository"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type PinMessageBiz struct {
	messageRepo messagerepo.Repository
}

func NewPinMessageBiz(messageRepo messagerepo.Repository) *PinMessageBiz {
	return &PinMessageBiz{messageRepo: messageRepo}
}

func (biz *PinMessageBiz) Pin(ctx context.Context, requesterId string, data *messagemdl.Message) error {
	log.Debug().Str("requesterId", requesterId).Any("data", data).Msg("pin message")

	data.Process()

	senderFilter, err := common.GetIdFilter(*data.SenderId)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid sender id"))
	}

	sender, err := biz.messageRepo.FindUser(ctx, senderFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find sender"))
	}

	if sender == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.Wrap(err, "sender not found"))
	}

	data.Sender = sender

	err = biz.messageRepo.CreateMessage(ctx, data)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not pin message"))
	}
	return nil
}
