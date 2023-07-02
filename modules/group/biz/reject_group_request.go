package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/pkg/errors"
)

type RejectGroupRequestBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.NotificationRepository
}

func NewRejectGroupRequestBiz(groupRepo grouprepo.Repository, notification notirepo.NotificationRepository) *RejectGroupRequestBiz {
	return &RejectGroupRequestBiz{groupRepo: groupRepo, notification: notification}
}

// RejectRequest send a group invitation request to user.
func (biz *RejectGroupRequestBiz) RejectRequest(ctx context.Context, requesterId string, groupId string) error {
	// Find exists request
	requesterFilter := requeststore.GetRequestReceiverIdFilter(requesterId)
	groupFilter := requeststore.GetRequestGroupIdFilter(groupId)
	ft := common.GetAndFilter(requesterFilter, groupFilter)
	err := biz.groupRepo.DeleteRequest(ctx, ft)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not delete request"))
	}
	return nil
}
