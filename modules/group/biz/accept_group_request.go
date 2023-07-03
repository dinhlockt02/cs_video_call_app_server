package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type AcceptGroupRequestBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.INotificationService
}

func NewAcceptGroupRequestBiz(groupRepo grouprepo.Repository,
	notification notirepo.INotificationService) *AcceptGroupRequestBiz {
	return &AcceptGroupRequestBiz{groupRepo: groupRepo, notification: notification}
}

// AcceptRequest send a group invitation request to user.
func (biz *AcceptGroupRequestBiz) AcceptRequest(ctx context.Context, requesterId string, groupId string) error {
	log.Debug().Str("requesterId", requesterId).Str("groupId", groupId).Msg("accept request")
	// Find exists request
	requestFilter := common.GetAndFilter(
		requeststore.GetRequestReceiverIdFilter(requesterId),
		requeststore.GetRequestGroupIdFilter(groupId),
	)
	existedRequest, err := biz.groupRepo.FindRequest(ctx, requestFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find existed request"))
	}
	if existedRequest == nil {
		return common.ErrInvalidRequest(errors.New(friendmodel.RequestNotFound))
	}

	// Find requester
	requesterFilter, err := common.GetIdFilter(requesterId)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid requester id"))
	}
	requester, err := biz.groupRepo.FindUser(ctx, requesterFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find requester"))
	}
	if requester == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.New(groupmdl.RequesterNotFound))
	}

	// Find Group
	groupFilter, err := common.GetIdFilter(groupId)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid group id"))
	}
	group, err := biz.groupRepo.FindGroup(ctx, groupFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find group"))
	}
	if group == nil {
		return common.ErrEntityNotFound(common.GroupEntity, errors.New(groupmdl.GroupNotFound))
	}

	// Update Requester
	requester.Groups = append(requester.Groups, groupId)

	err = biz.groupRepo.UpdateUser(ctx, requesterFilter, requester)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not update requester"))
	}

	// Update Group
	group.Members = append(group.Members, requesterId)
	err = biz.groupRepo.UpdateGroup(ctx, groupFilter, group)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not update group"))
	}

	// Delete request
	err = biz.groupRepo.DeleteRequest(ctx, requestFilter)
	if err != nil {
		return err
	}
	// TODO: send push notification new member joined
	return nil
}
