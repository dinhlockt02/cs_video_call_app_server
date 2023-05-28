package callrepo

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
)

func (repo *callRepository) UpdateCall(ctx context.Context, filter map[string]interface{}, data *callmdl.Call) error {
	return repo.callStore.Update(ctx, filter, data)
}
