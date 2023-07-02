package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type getGroupMembersBiz struct {
	groupRepo grouprepo.Repository
}

func NewGetGroupMembersBiz(groupRepo grouprepo.Repository) *getGroupMembersBiz {
	return &getGroupMembersBiz{groupRepo: groupRepo}
}

func (biz *getGroupMembersBiz) GetGroupUsers(ctx context.Context, userIds ...string) ([]groupmdl.User, error) {
	log.Debug().Any("userIds", userIds).Msg("get group users")
	users, err := biz.groupRepo.FindUsers(ctx, groupstore.GetUserIdInIdListFilter(userIds...))
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not find users"))
	}
	return users, nil
}
