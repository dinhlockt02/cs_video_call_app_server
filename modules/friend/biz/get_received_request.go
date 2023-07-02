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

type GetReceivedRequestBiz struct {
	friendRepo friendrepo.Repository
}

func NewGetReceivedRequestBiz(friendRepo friendrepo.Repository) *GetReceivedRequestBiz {
	return &GetReceivedRequestBiz{
		friendRepo: friendRepo,
	}
}

func (biz *GetReceivedRequestBiz) GetReceivedRequest(ctx context.Context, receiverId string) ([]requestmdl.Request, error) {
	log.Debug().Str("receiverId", receiverId).Msg("get received request")
	requests, err := biz.friendRepo.FindRequests(ctx, common.GetAndFilter(
		requeststore.GetRequestReceiverIdFilter(receiverId),
		requeststore.GetTypeFilterFilter(false),
	))
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not find received requests"))
	}

	return requests, nil
}
