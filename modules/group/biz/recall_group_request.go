package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
)

type recallGroupRequestBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.NotificationRepository
}

func NewRecallGroupRequestBiz(groupRepo grouprepo.Repository) *recallGroupRequestBiz {
	return &recallGroupRequestBiz{groupRepo: groupRepo}
}

// RecallRequest send a group invitation request to user.
func (biz *recallGroupRequestBiz) RecallRequest(ctx context.Context, requester string, user string, groupId string) error {
	// Find exists request
	senderFilter := requeststore.GetRequestSenderIdFilter(requester)
	receiverFilter := requeststore.GetRequestReceiverIdFilter(user)
	groupFilter := requeststore.GetRequestGroupIdFilter(groupId)
	ft := common.GetAndFilter(senderFilter, receiverFilter, groupFilter)
	existedRequest, err := biz.groupRepo.FindRequest(ctx, ft)
	if err != nil {
		return err
	}
	if existedRequest == nil {
		return common.ErrInvalidRequest(friendmodel.ErrRequestNotFound)
	}

	// Delete request
	filter := make(map[string]interface{})
	err = common.AddIdFilter(filter, *existedRequest.Id)
	err = biz.groupRepo.DeleteRequest(ctx, filter)
	if err != nil {
		return err
	}
	// TODO: send push notification new member joined
	return nil
}
