package callrepo

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	callstore "github.com/dinhlockt02/cs_video_call_app_server/modules/call/store"
	userstore "github.com/dinhlockt02/cs_video_call_app_server/modules/user/store"
)

type Repository interface {
	FindUser(
		ctx context.Context,
		filter map[string]interface{},
	) (*callmdl.User, error)
	CreateCall(ctx context.Context, call *callmdl.Call) error
	FindCall(ctx context.Context, filter map[string]interface{}) (*callmdl.Call, error)
	UpdateCall(ctx context.Context, filter map[string]interface{}, data *callmdl.UpdateCall) error
	ListCalls(ctx context.Context, filter map[string]interface{}) ([]callmdl.Call, error)
	RemoveOwner(ctx context.Context, filter map[string]interface{}, owner string) error
}

type CallRepository struct {
	userStore userstore.Store
	callStore callstore.Store
}

func NewCallRepository(
	userStore userstore.Store,
	callStore callstore.Store,
) *CallRepository {
	return &CallRepository{
		userStore: userStore,
		callStore: callStore,
	}
}
