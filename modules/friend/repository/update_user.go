package friendrepo

import (
	"context"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
)

// UpdateUser updates user that matches filter.
func (repo *FriendRepository) UpdateUser(
	ctx context.Context,
	filter map[string]interface{},
	user *friendmodel.User,
) error {
	return repo.friendstore.UpdateUser(ctx, filter, user)
}
