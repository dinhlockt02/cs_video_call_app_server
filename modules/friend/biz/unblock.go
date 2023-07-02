package friendbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type UnblockFriendStore interface {
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
	UpdateUser(ctx context.Context, filter map[string]interface{}, updatedUser *friendmodel.User) error
}

type UnblockBiz struct {
	friendStore UnfriendFriendStore
}

func NewUnblockBiz(friendStore UnfriendFriendStore) *UnblockBiz {
	return &UnblockBiz{
		friendStore: friendStore,
	}
}

func (biz *UnblockBiz) Unblock(ctx context.Context, userId string, blockedId string) error {
	log.Debug().Str("userId", userId).Str("blockedId", blockedId).Msg("unblock user")
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

	for i := range user.BlockedUser {
		if user.BlockedUser[i] == blockedId {

			user.BlockedUser = append(user.BlockedUser[:i], user.BlockedUser[i+1:]...)

			err = biz.friendStore.UpdateUser(ctx, userFilter, user)
			if err != nil {
				return errors.Wrap(err, "can not update user")
			}
			break
		}
	}
	return nil
}
