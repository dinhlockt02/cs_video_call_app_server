package friendbiz

import (
	"context"
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FindFriendFriendStore interface {
	FindFriend(ctx context.Context, filter map[string]interface{}) ([]friendmodel.FriendUser, error)
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
}

type findFriendBiz struct {
	friendStore FindFriendFriendStore
}

func NewFindFriendBiz(friendStore FindFriendFriendStore) *findFriendBiz {
	return &findFriendBiz{friendStore: friendStore}
}

func (biz *findFriendBiz) FindFriend(ctx context.Context, filter map[string]interface{}, friendFilter map[string]interface{}) ([]friendmodel.FriendUser, error) {
	user, err := biz.friendStore.FindUser(ctx, filter)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, common.ErrEntityNotFound("User", errors.New("user not found"))
	}

	var ids = make([]interface{}, 0, len(user.Friends))
	for _, friend := range user.Friends {
		id, _ := primitive.ObjectIDFromHex(friend)
		ids = append(ids, id)
	}
	friends, err := biz.friendStore.FindFriend(ctx,
		common.GetAndFilter(
			common.GetInFilter("_id", ids...),
			friendFilter,
		),
	)
	if err != nil {
		return nil, err
	}
	return friends, nil
}
