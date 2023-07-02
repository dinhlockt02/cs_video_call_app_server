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

type ListGroupBiz struct {
	groupRepo grouprepo.Repository
}

func NewListGroupBiz(groupRepo grouprepo.Repository) *ListGroupBiz {
	return &ListGroupBiz{groupRepo: groupRepo}
}

func (biz *ListGroupBiz) List(ctx context.Context,
	requesterId string, groupFilter map[string]interface{}) ([]groupmdl.Group, error) {
	log.Debug().Str("requesterId", requesterId).Any("groupFilter", groupFilter).Msg("leave")

	filter, err := common.GetIdFilter(requesterId)
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "invalid requester id"))
	}

	user, err := biz.groupRepo.FindUser(ctx, filter)
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not find user"))
	}

	filter = groupstore.GetGroupIdInIdListFilter(user.Groups...)

	groups, err := biz.groupRepo.List(ctx, common.GetAndFilter(filter, groupFilter))
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not find groups"))
	}
	return groups, nil
}
