package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	"github.com/pkg/errors"
)

type updateGroupBiz struct {
	groupRepo grouprepo.Repository
}

func NewUpdateGroupBiz(groupRepo grouprepo.Repository) *updateGroupBiz {
	return &updateGroupBiz{groupRepo: groupRepo}
}

func (biz *updateGroupBiz) Update(ctx context.Context, filter map[string]interface{}, data *groupmdl.Group) error {
	data.Members = nil

	if err := data.Process(); err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid update data"))
	}

	if err := biz.groupRepo.UpdateGroup(ctx, filter, data); err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "can not update group"))
	}

	return nil
}
