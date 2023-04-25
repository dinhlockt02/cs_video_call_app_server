package friendbiz

import (
	"context"
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlockFriendStore interface {
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
	UpdateUser(ctx context.Context, filter map[string]interface{}, updatedUser *friendmodel.User) error
}

type blockBiz struct {
	friendStore BlockFriendStore
}

func NewBlockBiz(friendStore UnfriendFriendStore) *blockBiz {
	return &blockBiz{
		friendStore: friendStore,
	}
}

func (biz *blockBiz) Block(ctx context.Context, userId string, blockedId string) error {
	id, _ := primitive.ObjectIDFromHex(userId)
	user, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})

	if err != nil {
		return err
	}
	if user == nil {
		return common.ErrEntityNotFound("User", errors.New("user not found"))
	}
	for i := range user.Friends {
		if user.Friends[i] == blockedId {
			user.Friends = append(user.Friends[:i], user.Friends[i+1:]...)
			break
		}
	}
	for i := range user.BlockedUser {
		if user.BlockedUser[i] == blockedId {
			return nil
		}
	}
	user.BlockedUser = append(user.BlockedUser, blockedId)
	err = biz.friendStore.UpdateUser(ctx, map[string]interface{}{
		"_id": id,
	}, user)
	if err != nil {
		return err
	}

	id, _ = primitive.ObjectIDFromHex(blockedId)
	blockedUser, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})
	if err != nil {
		return err
	}
	if blockedUser == nil {
		return common.ErrEntityNotFound("User", errors.New("receiver not found"))
	}
	for i := range blockedUser.Friends {
		if blockedUser.Friends[i] == userId {
			blockedUser.Friends = append(blockedUser.Friends[:i], blockedUser.Friends[i+1:]...)
			err = biz.friendStore.UpdateUser(ctx, map[string]interface{}{
				"_id": id,
			}, blockedUser)
			if err != nil {
				return err
			}
			break
		}
	}
	return nil
}
