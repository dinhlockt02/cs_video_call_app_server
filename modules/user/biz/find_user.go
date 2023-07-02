package userbiz

import (
	"context"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	"github.com/rs/zerolog/log"
)

type FindUserUserRepository interface {
	FindUser(ctx context.Context, requesterId string, filter map[string]interface{}) (*usermodel.User, error)
}
type FindUserBiz struct {
	userRepo FindUserUserRepository
}

func NewFindUserBiz(
	userRepo FindUserUserRepository,
) *FindUserBiz {
	return &FindUserBiz{
		userRepo: userRepo,
	}
}

func (biz *FindUserBiz) FindUser(ctx context.Context,
	requesterId string, filter map[string]interface{}) (*usermodel.User, error) {
	log.Debug().Str("requesterId", requesterId).Any("filter", filter).Msg("find user")
	return biz.userRepo.FindUser(ctx, requesterId, filter)
}
