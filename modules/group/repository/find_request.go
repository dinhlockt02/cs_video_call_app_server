package grouprepo

import (
	"context"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
)

// FindRequest returns the group request between sender and receiver
// If the request does not exist, it returns nil, nil
func (repo *groupRepository) FindRequest(
	ctx context.Context,
	filter map[string]interface{},
) (*requestmdl.Request, error) {

	existedRequest, err := repo.requestStore.FindRequest(ctx, filter)
	if err != nil {
		return nil, err
	}
	return existedRequest, nil
}
