package friendbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	friendrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/repository"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type RecallRequestBiz struct {
	friendRepo friendrepo.Repository
}

func NewRecallRequestBiz(friendRepo friendrepo.Repository) *RecallRequestBiz {
	return &RecallRequestBiz{
		friendRepo: friendRepo,
	}
}

func (biz *RecallRequestBiz) RecallRequest(ctx context.Context, senderId string, receiverId string) error {
	log.Debug().Str("senderId", senderId).Str("receiverId", receiverId).Msg("get sent request")
	existedRequest, err := biz.friendRepo.FindRequest(ctx, senderId, receiverId)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find request"))
	}
	if existedRequest == nil {
		return common.ErrInvalidRequest(errors.New(friendmodel.RequestNotFound))
	}
	filter, err := common.GetIdFilter(*existedRequest.Id)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "invalid request id: "+*existedRequest.Id))
	}
	err = biz.friendRepo.DeleteRequest(ctx, filter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not delete request"))
	}
	return nil
}
