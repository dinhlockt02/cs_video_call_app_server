package grouprepo

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
)

func (repo *groupRepository) FindUsers(ctx context.Context, filter map[string]interface{}) ([]groupmdl.User, error) {
	return repo.groupStore.FindUsers(ctx, filter)
}
