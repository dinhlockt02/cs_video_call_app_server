package groupbiz

import (
	"context"
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
)

type createGroupBiz struct {
	groupRepo grouprepo.Repository
}

func NewCreateGroupBiz(groupRepo grouprepo.Repository) *createGroupBiz {
	return &createGroupBiz{groupRepo: groupRepo}
}

// Create creates a group and add requester as a member.
func (biz *createGroupBiz) Create(ctx context.Context, requester string, data *groupmdl.Group) error {
	data.Members = []string{requester}

	if err := data.Process(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	userFilter := make(map[string]interface{})
	_ = common.AddIdFilter(userFilter, requester)
	user, err := biz.groupRepo.FindUser(ctx, userFilter)
	if err != nil {
		return err
	}
	if user == nil {
		return common.ErrEntityNotFound("User", errors.New("user not found"))
	}

	if err = biz.groupRepo.CreateGroup(ctx, data); err != nil {
		return err
	}

	user.Groups = append(user.Groups, *data.Id)

	if err = biz.groupRepo.UpdateUser(ctx, userFilter, user); err != nil {
		return err
	}

	return nil
}
