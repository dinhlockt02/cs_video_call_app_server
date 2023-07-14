package callbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	callrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/call/repository"
	callstore "github.com/dinhlockt02/cs_video_call_app_server/modules/call/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type DeleteCallsBiz struct {
	callRepo callrepo.Repository
}

func NewDeleteCallsBiz(
	callRepo callrepo.Repository,
) *DeleteCallsBiz {
	return &DeleteCallsBiz{
		callRepo: callRepo,
	}
}

func (biz *DeleteCallsBiz) DeleteCalls(ctx context.Context, requesterId string, filter map[string]interface{}) error {
	log.Debug().Str("requesterId", requesterId).Msg("can not delete calls")

	err := biz.callRepo.RemoveOwner(ctx, common.GetAndFilter(
		callstore.GetCallOwnerFilter(requesterId),
		filter), requesterId)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not delete calls"))
	}
	return nil
}
