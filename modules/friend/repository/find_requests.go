package friendrepo

import (
	"context"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
)

// FindRequests returns the friend request between sender and receiver
// If the request does not exist, it returns nil, nil.
func (repo *FriendRepository) FindRequests(
	ctx context.Context,
	filter map[string]interface{},
) ([]requestmdl.Request, error) {
	return repo.requestStore.FindRequests(ctx, filter)
}
