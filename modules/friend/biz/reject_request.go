package friendbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/pubsub"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	friendrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/repository"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type RejectRequestBiz struct {
	friendRepo friendrepo.Repository
	ps         pubsub.PubSub
}

func NewRejectRequestBiz(friendRepo friendrepo.Repository, ps pubsub.PubSub) *RejectRequestBiz {
	return &RejectRequestBiz{
		friendRepo: friendRepo,
		ps:         ps,
	}
}

func (biz *RejectRequestBiz) RejectRequest(ctx context.Context, senderId string, receiverId string) error {
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
	err = biz.ps.Publish(ctx, common.TopicRequestDeleted, *existedRequest.Id)
	if err != nil {
		log.Error().Stack().Err(err).Msg("can not publish event")
	}
	return nil
}
