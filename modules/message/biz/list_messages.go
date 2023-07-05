package messagebiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	messagemdl "github.com/dinhlockt02/cs_video_call_app_server/modules/message/model"
	messagerepo "github.com/dinhlockt02/cs_video_call_app_server/modules/message/repository"
	messagestore "github.com/dinhlockt02/cs_video_call_app_server/modules/message/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type ListMessagesBiz struct {
	messageRepo messagerepo.Repository
}

func NewListMessagesBiz(messageRepo messagerepo.Repository) *ListMessagesBiz {
	return &ListMessagesBiz{messageRepo: messageRepo}
}

func (biz *ListMessagesBiz) List(ctx context.Context, requesterId, groupId string) ([]messagemdl.Message, error) {
	log.Debug().Str("requesterId", requesterId).Str("groupId", groupId).Msg("list messages")
	messageFilter := messagestore.GetGroupIdFilter(groupId)

	messages, err := biz.messageRepo.ListMessages(ctx, messageFilter)
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not list messsages"))
	}
	return messages, nil
}
