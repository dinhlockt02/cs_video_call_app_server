package groupbiz

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
)

type getGroupMembersBiz struct {
	groupRepo grouprepo.Repository
}

func NewGetGroupMembersBiz(groupRepo grouprepo.Repository) *getGroupMembersBiz {
	return &getGroupMembersBiz{groupRepo: groupRepo}
}

func (biz *getGroupMembersBiz) GetGroupUsers(ctx context.Context, userIds ...string) ([]groupmdl.User, error) {
	return biz.groupRepo.FindUsers(ctx, groupstore.GetUserIdInIdListFilter(userIds...))
}
