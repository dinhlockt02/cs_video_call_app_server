package callrepo

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
)

func (repo *callRepository) FindUser(ctx context.Context, filter map[string]interface{}) (*callmdl.User, error) {
	user, err := repo.userStore.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	return &callmdl.User{
		Id:     *user.Id,
		Name:   user.Name,
		Avatar: user.Avatar,
	}, nil
}
