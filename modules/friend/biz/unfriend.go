package friendbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type UnfriendFriendStore interface {
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
	UpdateUser(ctx context.Context, filter map[string]interface{}, updatedUser *friendmodel.User) error
}

type UnfriendBiz struct {
	friendStore UnfriendFriendStore
}

func NewUnfriendBiz(friendStore UnfriendFriendStore) *UnfriendBiz {
	return &UnfriendBiz{
		friendStore: friendStore,
	}
}

func (biz *UnfriendBiz) Unfriend(ctx context.Context, userId string, friendId string) error {
	log.Debug().Str("userId", userId).Str("friendId", friendId).Msg("unfriend user")
	userFilter, err := common.GetIdFilter(userId)
	if err != nil {
		return common.ErrInvalidRequest(errors.New("invalid user id"))
	}
	user, err := biz.friendStore.FindUser(ctx, userFilter)

	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find user"))
	}
	if user == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.New(friendmodel.UserNotFound))
	}
	for i := range user.Friends {
		if user.Friends[i] == friendId {
			user.Friends = append(user.Friends[:i], user.Friends[i+1:]...)
			break
		}
	}

	friendFilter, err := common.GetIdFilter(friendId)
	if err != nil {
		return common.ErrInvalidRequest(errors.New("invalid friend id"))
	}
	friend, err := biz.friendStore.FindUser(ctx, friendFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find friend"))
	}
	if friend == nil {
		return common.ErrEntityNotFound("User", errors.New(friendmodel.UserNotFound))
	}
	for i := range friend.Friends {
		if friend.Friends[i] == userId {
			friend.Friends = append(friend.Friends[:i], friend.Friends[i+1:]...)
			break
		}
	}
	err = biz.friendStore.UpdateUser(ctx, friendFilter, friend)
	if err != nil {
		return err
	}
	err = biz.friendStore.UpdateUser(ctx, userFilter, user)
	if err != nil {
		return err
	}
	return nil
}
