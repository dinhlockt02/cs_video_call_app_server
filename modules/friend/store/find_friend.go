package friendstore

import (
	"context"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *MongoStore) FindFriends(ctx context.Context, filter map[string]interface{}) ([]friendmodel.FriendUser, error) {
	log.Debug().Any("filter", filter).Msg("find friends")
	var friends []friendmodel.FriendUser
	cur, err := s.database.Collection(friendmodel.FriendUser{}.CollectionName()).Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find friend")
	}
	if err = cur.All(ctx, &friends); err != nil {
		return nil, errors.Wrap(err, "can not decode friend data")
	}

	return friends, nil
}
