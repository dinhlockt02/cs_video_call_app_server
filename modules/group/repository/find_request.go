package grouprepo

import (
	"context"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
)

// FindRequest returns the group request between sender and receiver
// If the request does not exist, it returns nil, nil.
func (repo *GroupRepository) FindRequest(
	ctx context.Context,
	filter map[string]interface{},
) (*requestmdl.Request, error) {
	return repo.requestStore.FindRequest(ctx, filter)
}
