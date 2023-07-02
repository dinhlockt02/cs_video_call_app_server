package callrepo

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (repo *callRepository) FindUser(ctx context.Context, filter map[string]interface{}) (*callmdl.User, error) {
	log.Debug().Any("filter", filter).Msg("find user")
	user, err := repo.userStore.Find(ctx, filter)

	if err != nil {
		return nil, errors.Wrap(err, "can not find user")
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
