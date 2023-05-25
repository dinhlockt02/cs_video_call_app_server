package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
)

type createGroupBiz struct {
	groupStore GroupStore
}

func NewCreateGroupBiz(groupStore GroupStore) *createGroupBiz {
	return &createGroupBiz{groupStore: groupStore}
}

func (biz *createGroupBiz) Create(ctx context.Context, requester string, data *groupmdl.Group) error {
	data.Members = []string{requester}

	if err := data.Process(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := biz.groupStore.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
