package callbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	callrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/call/repository"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type ListCallsBiz struct {
	callRepo callrepo.Repository
}

func NewListCallsBiz(
	callRepo callrepo.Repository,
) *ListCallsBiz {
	return &ListCallsBiz{
		callRepo: callRepo,
	}
}

func (biz *ListCallsBiz) List(ctx context.Context, filter map[string]interface{}) ([]callmdl.Call, error) {
	log.Debug().Any("filter", filter).
		Msg("list calls")
	calls, err := biz.callRepo.ListCalls(ctx, filter)
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not list calls"))
	}
	return calls, nil
}
