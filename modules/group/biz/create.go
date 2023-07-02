package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type createGroupBiz struct {
	groupRepo grouprepo.Repository
}

func NewCreateGroupBiz(groupRepo grouprepo.Repository) *createGroupBiz {
	return &createGroupBiz{groupRepo: groupRepo}
}

// Create creates a group and add requester as a member.
func (biz *createGroupBiz) Create(ctx context.Context, requesterId string, data *groupmdl.Group) error {
	log.Debug().Str("requesterId", requesterId).Any("data", data).Msg("create group")
	data.Members = []string{requesterId}

	if err := data.Process(); err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid group data"))
	}

	requesterFilter, err := common.GetIdFilter(requesterId)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "invalid requester id"))
	}
	requester, err := biz.groupRepo.FindUser(ctx, requesterFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find user"))
	}
	if requester == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.New(groupmdl.RequesterNotFound))
	}

	if err = biz.groupRepo.CreateGroup(ctx, data); err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not create group"))
	}

	requester.Groups = append(requester.Groups, *data.Id)

	if err = biz.groupRepo.UpdateUser(ctx, requesterFilter, requester); err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not update user"))
	}

	return nil
}
