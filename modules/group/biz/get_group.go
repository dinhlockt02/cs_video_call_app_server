package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type GetGroupBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.Repository
}

func NewGetGroupBiz(groupRepo grouprepo.Repository, notification notirepo.Repository) *GetGroupBiz {
	return &GetGroupBiz{groupRepo: groupRepo, notification: notification}
}

// GetById returns a group by id.
func (biz *GetGroupBiz) GetById(ctx context.Context, groupId string) (*groupmdl.Group, error) {
	log.Debug().Str("groupId", groupId).Msg("get group by id")
	// Find Group
	filter, err := common.GetIdFilter(groupId)
	if err != nil {
		return nil, common.ErrInvalidRequest(errors.Wrap(err, "invalid group id"))
	}
	group, err := biz.groupRepo.FindGroup(ctx, filter)
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not find group"))
	}
	if group == nil {
		return nil, common.ErrEntityNotFound(common.GroupEntity, errors.New(groupmdl.GroupNotFound))
	}

	return group, nil
}
