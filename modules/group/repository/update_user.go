package grouprepo

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
)

func (repo *groupRepository) UpdateUser(
	ctx context.Context,
	filter map[string]interface{},
	updatedUser *groupmdl.User,
) error {
	return repo.groupStore.UpdateUser(ctx, filter, updatedUser)
}
