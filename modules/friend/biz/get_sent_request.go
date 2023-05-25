package friendbiz

import (
	"context"
	friendrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/repository"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
)

type getSentRequestBiz struct {
	friendRepo friendrepo.Repository
}

func NewGetSentRequestBiz(friendRepo friendrepo.Repository) *getSentRequestBiz {
	return &getSentRequestBiz{
		friendRepo: friendRepo,
	}
}

func (biz *getSentRequestBiz) GetSentRequest(ctx context.Context, senderId string) ([]requestmdl.Request, error) {
	requests, err := biz.friendRepo.FindRequests(ctx, requeststore.GetRequestSenderIdFilter(senderId))
	if err != nil {
		return nil, err
	}

	return requests, nil
}
