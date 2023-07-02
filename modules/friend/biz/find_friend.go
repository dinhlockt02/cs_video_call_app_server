package friendbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type FindFriendFriendStore interface {
	FindFriends(ctx context.Context, filter map[string]interface{}) ([]friendmodel.FriendUser, error)
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
}

type FindFriendBiz struct {
	friendStore FindFriendFriendStore
}

func NewFindFriendBiz(friendStore FindFriendFriendStore) *FindFriendBiz {
	return &FindFriendBiz{friendStore: friendStore}
}

func (biz *FindFriendBiz) FindFriend(ctx context.Context, filter map[string]interface{},
	friendFilter map[string]interface{}) ([]friendmodel.FriendUser, error) {
	log.Debug().Any("filter", filter).Any("friendFilter", friendFilter).Msg("find friends")

	user, err := biz.friendStore.FindUser(ctx, filter)

	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not find user"))
	}

	if user == nil {
		return nil, common.ErrEntityNotFound(common.UserEntity, errors.New(friendmodel.UserNotFound))
	}

	var ids = make([]interface{}, 0, len(user.Friends))
	for _, friend := range user.Friends {
		id, err := common.ToObjectId(friend)
		if err != nil {
			log.Error().Stack().Err(err).Str("friendId", friend).Msg("database crashed: invalid id injected.")
			continue
		}
		ids = append(ids, id)
	}
	friends, err := biz.friendStore.FindFriends(ctx,
		common.GetAndFilter(
			common.GetInFilter("_id", ids...),
			friendFilter,
		),
	)
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not find friend"))
	}
	return friends, nil
}
