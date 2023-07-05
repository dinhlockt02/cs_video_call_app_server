package messagerepo

import (
	"context"
	messagemdl "github.com/dinhlockt02/cs_video_call_app_server/modules/message/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (repo *MessageRepository) FindUser(ctx context.Context, filter map[string]interface{}) (*messagemdl.User, error) {
	log.Debug().Any("filter", filter).Msg("find user")
	user, err := repo.userStore.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find user")
	}
	if user == nil {
		return nil, errors.Wrap(err, "user not found")
	}
	return &messagemdl.User{
		Id:     *user.Id,
		Name:   user.Name,
		Avatar: user.Avatar,
	}, nil
}
