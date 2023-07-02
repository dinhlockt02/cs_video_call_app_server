package grouprepo

import (
	"context"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
)

func (repo *GroupRepository) CreateRequest(
	ctx context.Context,
	request *requestmdl.Request,
) error {
	return repo.requestStore.CreateRequest(ctx, request)
}
