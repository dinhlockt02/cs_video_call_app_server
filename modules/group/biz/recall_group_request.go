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

type RecallGroupRequestBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.INotificationService
	ps           pubsub.PubSub
}

func NewRecallGroupRequestBiz(groupRepo grouprepo.Repository,
	notification notirepo.INotificationService,
	ps pubsub.PubSub) *RecallGroupRequestBiz {
	return &RecallGroupRequestBiz{groupRepo: groupRepo, notification: notification, ps: ps}
}

func (biz *RecallGroupRequestBiz) RecallRequest(ctx context.Context, requesterId string, groupId string) error {
	// Find exists request
	requesterFilter := requeststore.GetRequestSenderIdFilter(requesterId)
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
