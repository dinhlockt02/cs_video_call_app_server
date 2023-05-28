package callrepo

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
)

func (repo *callRepository) CreateCall(ctx context.Context, call *callmdl.Call) error {
	err := repo.callStore.Create(ctx, call)

	if err != nil {
		return err
	}

	return nil
}
