package grouprepo

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
)

func (repo *GroupRepository) UpdateGroup(
	ctx context.Context,
	filter map[string]interface{},
	updatedGroup *groupmdl.Group,
) error {
	return repo.groupStore.UpdateGroup(ctx, filter, updatedGroup)
}
