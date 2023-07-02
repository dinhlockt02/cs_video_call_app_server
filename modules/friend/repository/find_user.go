package friendrepo

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// FindUser is a method of finding a user.
func (repo *FriendRepository) FindUser(
	ctx context.Context,
	filter map[string]interface{},
	options ...FindUserOption,
) (*friendmodel.User, error) {
	log.Debug().Any("filter", filter).Any("options", options).Msg("find a user")
	user, err := repo.friendstore.FindUser(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find user")
	}

	if user == nil {
		return nil, common.ErrEntityNotFound(common.UserEntity, errors.New(friendmodel.UserNotFound))
	}

	for _, option := range options {
		err = option(ctx, repo, user)
		if err != nil {
			return nil, errors.Wrap(err, "failed to apply option")
		}
	}

	return user, nil
}

type FindUserOption func(ctx context.Context, repo *FriendRepository, u *friendmodel.User) error

// WithRelation requires requesterId, which is the id of the requester,
//
// It will use the requesterId to query the relation of user with requester
// The relation includes
// [friendmodel.Self], [friendmodel.Friend], [friendmodel.Blocked],
// [friendmodel.Non], [friendmodel.Sent], [friendmodel.Received].
func WithRelation(requesterId string) FindUserOption {
	return func(ctx context.Context, repo *FriendRepository, user *friendmodel.User) error {
		log.Debug().Msg("add relation to user model")
		if *user.Id == requesterId {
			user.Relation = friendmodel.Self
			return nil
		}

		for _, friendId := range user.Friends {
			if friendId == requesterId {
				user.Relation = friendmodel.Friend
				return nil
			}
		}

		filter, err := common.GetIdFilter(requesterId)

		if err != nil {
			return errors.Wrap(err, "invalid requester id "+requesterId)
		}

		requester, err := repo.friendstore.FindUser(ctx, filter)

		if err != nil {
			return errors.Wrap(err, "can not find requester")
		}

		if requester == nil {
			return common.ErrEntityNotFound(common.UserEntity, errors.New(friendmodel.UserNotFound))
		}

		for _, blockedId := range requester.BlockedUser {
			if *user.Id == blockedId {
				user.Relation = friendmodel.Blocked
				return nil
			}
		}

		filter = common.GetOrFilter(
			common.GetAndFilter(
				requeststore.GetRequestSenderIdFilter(requesterId),
				requeststore.GetRequestReceiverIdFilter(*user.Id),
			),
			common.GetAndFilter(
				requeststore.GetRequestSenderIdFilter(*user.Id),
				requeststore.GetRequestReceiverIdFilter(requesterId),
			),
		)

		request, err := repo.requestStore.FindRequest(ctx, filter)
		if err != nil {
			return errors.Wrap(err, "can not find request")
		}

		if request == nil {
			user.Relation = friendmodel.Non
			return nil
		}

		if request.Sender.Id == requesterId {
			user.Relation = friendmodel.Sent
			return nil
		}

		user.Relation = friendmodel.Received
		return nil
	}
}
