package friendbiz

import (
	"context"
	"github.com/rs/zerolog/log"

	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	"github.com/pkg/errors"
)

type BlockFriendStore interface {
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
	UpdateUser(ctx context.Context, filter map[string]interface{}, updatedUser *friendmodel.User) error
}

type BlockBiz struct {
	friendStore BlockFriendStore
}

func NewBlockBiz(friendStore UnfriendFriendStore) *BlockBiz {
	return &BlockBiz{
		friendStore: friendStore,
	}
}

func (biz *BlockBiz) Block(ctx context.Context, userId string, blockedId string) error {
	log.Debug().Str("userId", userId).Str("blockedId", blockedId).Msg("block user")
	userIdFilter, err := common.GetIdFilter(userId)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid userId"))
	}
	user, err := biz.friendStore.FindUser(ctx, userIdFilter)

	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find user"))
	}
	if user == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.New(friendmodel.UserNotFound))
	}

	// Remove from friend list
	for i := range user.Friends {
		if user.Friends[i] == blockedId {
			user.Friends = append(user.Friends[:i], user.Friends[i+1:]...)
			break
		}
	}

	// Exit if user is blocked
	for i := range user.BlockedUser {
		if user.BlockedUser[i] == blockedId {
			return nil
		}
	}

	user.BlockedUser = append(user.BlockedUser, blockedId)

	blockedUserFilter, err := common.GetIdFilter(blockedId)
	if err != nil {
		return common.ErrInvalidRequest(errors.New("invalid blocked user id"))
	}
	blockedUser, err := biz.friendStore.FindUser(ctx, blockedUserFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find blocked user"))
	}
	if blockedUser == nil {
		return common.ErrEntityNotFound("User", errors.New(friendmodel.BlockedUserNotFound))
	}
	for i := range blockedUser.Friends {
		if blockedUser.Friends[i] == userId {
			blockedUser.Friends = append(blockedUser.Friends[:i], blockedUser.Friends[i+1:]...)
			break
		}
	}

	err = biz.friendStore.UpdateUser(ctx, userIdFilter, user)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not update user"))
	}

	err = biz.friendStore.UpdateUser(ctx, blockedUserFilter, blockedUser)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not update blocked user"))
	}
	return nil
}
