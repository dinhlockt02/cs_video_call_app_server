package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type RecallGroupRequestBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.INotificationService
}

func NewRecallGroupRequestBiz(groupRepo grouprepo.Repository,
	notification notirepo.INotificationService) *RecallGroupRequestBiz {
	return &RecallGroupRequestBiz{groupRepo: groupRepo, notification: notification}
}

// RecallRequest send a group invitation request to user.
func (biz *RecallGroupRequestBiz) RecallRequest(ctx context.Context,
	requesterId string, user string, groupId string) error {
	log.Debug().Str("requesterId", requesterId).Str("user", user).Str("groupId", groupId).Msg("recall request")
	// Find exists request
	senderFilter := requeststore.GetRequestSenderIdFilter(requesterId)
	receiverFilter := requeststore.GetRequestReceiverIdFilter(user)
	groupFilter := requeststore.GetRequestGroupIdFilter(groupId)
	ft := common.GetAndFilter(senderFilter, receiverFilter, groupFilter)
	err := biz.groupRepo.DeleteRequest(ctx, ft)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not delete request"))
	}
	return nil
}
