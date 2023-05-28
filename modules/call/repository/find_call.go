package callrepo

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
)

func (repo *callRepository) FindCall(ctx context.Context, filter map[string]interface{}) (*callmdl.Call, error) {
	return repo.callStore.FindCall(ctx, filter)
}
