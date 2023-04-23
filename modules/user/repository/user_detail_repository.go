package userrepo

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDetailUserStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error)
}

type UserDetailFriendStore interface {
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
}

type userDetailRepository struct {
	userStore   UserDetailUserStore
	friendStore UserDetailFriendStore
}

func NewUserRepository(userStore UserDetailUserStore,
	friendStore UserDetailFriendStore,
) *userDetailRepository {
	return &userDetailRepository{
		userStore:   userStore,
		friendStore: friendStore,
	}
}

func (r *userDetailRepository) GetUserDetail(ctx context.Context, userId string, otherId string) (*usermodel.UserDetail, error) {

	id, _ := primitive.ObjectIDFromHex(userId)
	user, err := r.userStore.Find(ctx, map[string]interface{}{
		"_id": id,
	})

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, common.ErrEntityNotFound("User", usermodel.ErrUserNotFound)
	}

	userDetail := usermodel.NewUserDetail(user)

	friendStoreUser, err := r.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})

	if err != nil {
		return nil, err
	}
	if friendStoreUser == nil {
		return nil, common.ErrEntityNotFound("User", usermodel.ErrUserNotFound)
	}

	userFriendMap := make(map[string]interface{}, len(friendStoreUser.Friends))

	for _, friend := range friendStoreUser.Friends {
		if friend == otherId {
			userDetail.IsFriend = true
		}
		userFriendMap[friend] = struct{}{}
	}

	id, _ = primitive.ObjectIDFromHex(otherId)
	friendStoreUser, err = r.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})

	if err != nil {
		return nil, err
	}
	if friendStoreUser == nil {
		return nil, common.ErrEntityNotFound("User", usermodel.ErrUserNotFound)
	}

	commonFriend := make([]string, 0)
	for _, friend := range friendStoreUser.Friends {
		if _, ok := userFriendMap[friend]; ok {
			commonFriend = append(commonFriend, friend)
		}
	}
	userDetail.CommonFriend = commonFriend
	userDetail.CommonFriendCount = len(commonFriend)

	return userDetail, nil
}
