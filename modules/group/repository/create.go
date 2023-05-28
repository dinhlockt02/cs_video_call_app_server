package grouprepo

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
)

func (repo *groupRepository) CreateGroup(ctx context.Context, group *groupmdl.Group) error {
	return repo.groupStore.Create(ctx, group)
}
