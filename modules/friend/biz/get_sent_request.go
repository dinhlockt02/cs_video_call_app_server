package friendbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/repository"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type GetSentRequestBiz struct {
	friendRepo friendrepo.Repository
}

func NewGetSentRequestBiz(friendRepo friendrepo.Repository) *GetSentRequestBiz {
	return &GetSentRequestBiz{
		friendRepo: friendRepo,
	}
}

func (biz *GetSentRequestBiz) GetSentRequest(ctx context.Context, senderId string) ([]requestmdl.Request, error) {
	log.Debug().Str("senderId", senderId).Msg("get sent request")
	requests, err := biz.friendRepo.FindRequests(ctx, common.GetAndFilter(
		requeststore.GetRequestSenderIdFilter(senderId),
		requeststore.GetTypeFilterFilter(false),
	))
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not find sent requests"))
	}

	return requests, nil
}
