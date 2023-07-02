package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type leaveGroupBiz struct {
	groupRepo grouprepo.Repository
}

func NewLeaveGroupBiz(groupRepo grouprepo.Repository) *leaveGroupBiz {
	return &leaveGroupBiz{groupRepo: groupRepo}
}

func (biz *leaveGroupBiz) Leave(ctx context.Context, userFilter map[string]interface{}, groupFilter map[string]interface{}) error {
	log.Debug().Any("userFilter", userFilter).Any("groupFilter", groupFilter).Msg("leave")
	user, err := biz.groupRepo.FindUser(ctx, userFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find user"))
	}
	if user == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.New(groupmdl.UserNotFound))
	}
	group, err := biz.groupRepo.FindGroup(ctx, groupFilter)
	if err != nil {
		return err
	}
	if group == nil {
		return common.ErrEntityNotFound(common.GroupEntity, errors.New(groupmdl.GroupNotFound))
	}
	for i, member := range group.Members {
		if member == *user.Id {
			group.Members = append(group.Members[:i], group.Members[i+1:]...)
			break
		}
	}

	for i, groupId := range user.Groups {
		if groupId == *group.Id {
			user.Groups = append(user.Groups[:i], user.Groups[i+1:]...)
			break
		}
	}

	if len(group.Members) > 0 {
		err = biz.groupRepo.UpdateGroup(ctx, groupFilter, group)
		if err != nil {
			return err
		}
	} else {
		err = biz.groupRepo.DeleteOne(ctx, groupFilter)
		if err != nil {
			return err
		}
	}
	err = biz.groupRepo.UpdateUser(ctx, userFilter, user)
	if err != nil {
		return err
	}
	return nil
}
