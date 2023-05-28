package grouprepo

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
)

func (repo *groupRepository) List(ctx context.Context, groupFilter map[string]interface{}) ([]groupmdl.Group, error) {
	return repo.groupStore.List(ctx, groupFilter)
}
