package userbiz

import (
	"context"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
)

type UserDetailRepo interface {
	GetUserDetail(ctx context.Context, userId string, otherId string) (*usermodel.UserDetail, error)
}

type userDetailBiz struct {
	repo UserDetailRepo
}

func NewUserDetailBiz(repo UserDetailRepo) *userDetailBiz {
	return &userDetailBiz{repo: repo}
}

func (biz *userDetailBiz) GetUserDetail(ctx context.Context, userId string, requesterId string) (*usermodel.UserDetail, error) {
	user, err := biz.repo.GetUserDetail(ctx, userId, requesterId)
	if userId == requesterId {
		user.CommonFriendCount = nil
		user.CommonFriend = nil
		user.IsFriend = nil
	}
	return user, err
}
