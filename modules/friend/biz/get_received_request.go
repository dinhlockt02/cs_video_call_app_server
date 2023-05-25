package friendbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/repository"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
)

type getReceivedRequestBiz struct {
	friendRepo friendrepo.Repository
}

func NewGetReceivedRequestBiz(friendRepo friendrepo.Repository) *getReceivedRequestBiz {
	return &getReceivedRequestBiz{
		friendRepo: friendRepo,
	}
}

func (biz *getReceivedRequestBiz) GetReceivedRequest(ctx context.Context, receiverId string) ([]requestmdl.Request, error) {
	requests, err := biz.friendRepo.FindRequests(ctx, common.GetAndFilter(
		requeststore.GetRequestReceiverIdFilter(receiverId),
		requeststore.GetTypeFilterFilter(false),
	))
	if err != nil {
		return nil, err
	}

	return requests, nil
}
