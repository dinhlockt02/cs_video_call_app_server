package groupbiz

import (
	"context"
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
)

type getGroupBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.NotificationRepository
}

func NewGetGroupBiz(groupRepo grouprepo.Repository, notification notirepo.NotificationRepository) *getGroupBiz {
	return &getGroupBiz{groupRepo: groupRepo, notification: notification}
}

// GetById returns a group by id.
func (biz *getGroupBiz) GetById(ctx context.Context, groupId string) (*groupmdl.Group, error) {
	// Find Group
	filter := make(map[string]interface{})
	err := common.AddIdFilter(filter, groupId)
	if err != nil {
		return nil, common.ErrInvalidRequest(errors.New("invalid group id"))
	}
	group, err := biz.groupRepo.FindGroup(ctx, filter)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, common.ErrEntityNotFound("Group", errors.New("group not found"))
	}

	return group, nil
}
