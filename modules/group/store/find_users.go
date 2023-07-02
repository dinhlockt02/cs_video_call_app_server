package groupstore

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *mongoStore) FindUsers(ctx context.Context, filter map[string]interface{}) ([]groupmdl.User, error) {
	log.Debug().Any("filter", filter).Msg("find users")
	var users []groupmdl.User
	cursor, err := s.database.Collection(groupmdl.User{}.CollectionName()).Find(ctx, filter)

	if err != nil {
		return nil, errors.Wrap(err, "can not find users")
	}

	if err = cursor.All(ctx, &users); err != nil {
		return nil, errors.Wrap(err, "can not decode users")
	}

	return users, nil
}
