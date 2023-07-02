package callrepo

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (repo *callRepository) FindCall(ctx context.Context, filter map[string]interface{}) (*callmdl.Call, error) {
	log.Debug().Any("filter", filter).Msg("find call")
	call, err := repo.callStore.FindCall(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find call")
	}
	return call, nil
}
