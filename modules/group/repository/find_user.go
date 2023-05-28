package grouprepo

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
)

func (repo *groupRepository) FindUser(ctx context.Context, filter map[string]interface{}) (*groupmdl.User, error) {
	return repo.groupStore.FindUser(ctx, filter)
}
