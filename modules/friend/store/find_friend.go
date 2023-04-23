package friendstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
)

func (s *mongoStore) FindFriend(ctx context.Context, filter map[string]interface{}) ([]friendmodel.FriendUser, error) {
	var friends []friendmodel.FriendUser
	cur, err := s.database.Collection(friendmodel.FriendUser{}.CollectionName()).Find(ctx, filter)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	if err = cur.All(ctx, &friends); err != nil {
		return nil, common.ErrInternal(err)
	}

	return friends, nil
}
