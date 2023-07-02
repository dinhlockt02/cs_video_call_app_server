package userrepo

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/repository"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type FindUserUserStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error)
}

type FindUserRepo struct {
	userStore  FindUserUserStore
	friendRepo friendrepo.Repository
}

func NewFindUserRepo(
	userStore FindUserUserStore,
	friendRepo friendrepo.Repository,
) *FindUserRepo {
	return &FindUserRepo{
		userStore:  userStore,
		friendRepo: friendRepo,
	}
}

func (repo *FindUserRepo) FindUser(ctx context.Context, requesterId string,
	filter map[string]interface{}) (*usermodel.User, error) {
	log.Debug().Str("requesterId", requesterId).Any("filter", filter).Msg("find user")
	fuser, err := repo.friendRepo.FindUser(ctx, filter, friendrepo.WithRelation(requesterId))

	if err != nil {
		return nil, errors.Wrap(err, "can not find user")
	}

	user, err := repo.userStore.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find user")
	}

	user.Relation = fuser.Relation

	filter, err = common.GetIdFilter(requesterId)
	if err != nil {
		return nil, errors.Wrap(err, "invalid requester id")
	}

	requester, err := repo.friendRepo.FindUser(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find requester")
	}

	requesterFriendMap := make(map[string]interface{}, len(requester.Friends))

	for _, friend := range requester.Friends {
		requesterFriendMap[friend] = struct{}{}
	}

	commonFriendCount := 0
	for _, friend := range fuser.Friends {
		if _, ok := requesterFriendMap[friend]; ok {
			commonFriendCount++
		}
	}

	user.CommonFriendCount = &commonFriendCount

	return user, nil
}
