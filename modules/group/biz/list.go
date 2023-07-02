package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
)

type listGroupBiz struct {
	groupRepo grouprepo.Repository
}

func NewListGroupBiz(groupRepo grouprepo.Repository) *listGroupBiz {
	return &listGroupBiz{groupRepo: groupRepo}
}

func (biz *listGroupBiz) List(ctx context.Context, requesterId string, groupFilter map[string]interface{}) ([]groupmdl.Group, error) {

	filter := make(map[string]interface{})
	_ = common.AddIdFilter(filter, requesterId)

	user, err := biz.groupRepo.FindUser(ctx, filter)
	if err != nil {
		return nil, err
	}

	filter = groupstore.GetGroupIdInIdListFilter(user.Groups...)

	return biz.groupRepo.List(ctx, common.GetAndFilter(filter, groupFilter))

}
