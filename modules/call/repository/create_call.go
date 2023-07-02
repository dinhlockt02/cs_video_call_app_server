package callrepo

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (repo *callRepository) CreateCall(ctx context.Context, call *callmdl.Call) error {
	log.Debug().Any("call", call).Msg("create call")
	err := repo.callStore.Create(ctx, call)
	if err != nil {
		return errors.Wrap(err, "can not create call")
	}

	return nil
}
