package callrepo

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (repo *CallRepository) UpdateCall(ctx context.Context, filter map[string]interface{}, data *callmdl.Call) error {
	log.Debug().Any("filter", filter).Any("data", data).Msg("update call")
	err := repo.callStore.Update(ctx, filter, data)
	if err != nil {
		return errors.Wrap(err, "can not update call")
	}
	return nil
}
