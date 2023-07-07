package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	"github.com/dinhlockt02/cs_video_call_app_server/components/pubsub"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type RejectGroupRequestBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.INotificationService
	ps           pubsub.PubSub
}

func NewRejectGroupRequestBiz(groupRepo grouprepo.Repository,
	notification notirepo.INotificationService,
	ps pubsub.PubSub) *RejectGroupRequestBiz {
	return &RejectGroupRequestBiz{groupRepo: groupRepo, notification: notification, ps: ps}
}

// RejectRequest send a group invitation request to user.
func (biz *RejectGroupRequestBiz) RejectRequest(ctx context.Context, requesterId string, groupId string) error {
	// Find exists request
	requesterFilter := requeststore.GetRequestReceiverIdFilter(requesterId)
	groupFilter := requeststore.GetRequestGroupIdFilter(groupId)
	ft := common.GetAndFilter(requesterFilter, groupFilter)
	existedRequest, err := biz.groupRepo.FindRequest(ctx, ft)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find existed request"))
	}
	if existedRequest == nil {
		return common.ErrInvalidRequest(errors.New(groupmdl.RequesterNotFound))
	}

	err = biz.groupRepo.DeleteRequest(ctx, ft)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not delete request"))
	}
	err = biz.ps.Publish(ctx, common.TopicRequestDeleted, *existedRequest.Id)
	if err != nil {
		log.Error().Stack().Err(err).Msg("can not publish event")
	}
	return nil
}
