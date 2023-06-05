package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
)

type getGroupRequestBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.NotificationRepository
}

func NewGetGroupRequestBiz(groupRepo grouprepo.Repository) *getGroupRequestBiz {
	return &getGroupRequestBiz{groupRepo: groupRepo}
}

// GetRequest send a group invitation request to user.
func (biz *getGroupRequestBiz) GetRequest(ctx context.Context, requesterId string, filter groupmdl.Filter) ([]requestmdl.Request, error) {

	var groupFilterFilter map[string]interface{}
	if filter == groupmdl.Sent {
		groupFilterFilter = requeststore.GetRequestSenderIdFilter(requesterId)
	} else {
		groupFilterFilter = requeststore.GetRequestReceiverIdFilter(requesterId)
	}

	ft := common.GetAndFilter(
		groupFilterFilter,
		requeststore.GetTypeFilterFilter(true),
	)

	requests, err := biz.groupRepo.FindRequests(ctx, ft)

	if err != nil {
		return nil, err
	}

	return requests, nil
}