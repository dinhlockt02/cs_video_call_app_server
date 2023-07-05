package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type SendGroupRequestBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.INotificationService
}

func NewSendGroupRequestBiz(groupRepo grouprepo.Repository,
	notification notirepo.INotificationService) *SendGroupRequestBiz {
	return &SendGroupRequestBiz{groupRepo: groupRepo, notification: notification}
}

// SendRequest send a group invitation request to user.
func (biz *SendGroupRequestBiz) SendRequest(ctx context.Context,
	requesterId string, user string, group *groupmdl.Group) error {
	// TODO: Allow send request only if requester is a member of group

	// Find exists request
	requestFilter := common.GetAndFilter(
		requeststore.GetRequestReceiverIdFilter(user),
		requeststore.GetRequestGroupIdFilter(*group.Id),
	)
	existedRequest, err := biz.groupRepo.FindRequest(ctx, requestFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find request"))
	}
	if existedRequest != nil {
		return common.ErrInvalidRequest(errors.New(friendmodel.RequestExists))
	}

	// Find requester
	requesterFilter, err := common.GetIdFilter(requesterId)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "invalid requester id"))
	}
	requester, err := biz.groupRepo.FindUser(ctx, requesterFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find requester"))
	}
	if requester == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.New("requester not found"))
	}

	// Find Receiver
	receiverFilter, err := common.GetIdFilter(user)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid receiver id"))
	}
	receiver, err := biz.groupRepo.FindUser(ctx, receiverFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find receiver"))
	}
	if receiver == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.New("receiver not found"))
	}

	senderRequestUser := requestmdl.RequestUser{
		Id:     requesterId,
		Name:   requester.Name,
		Avatar: requester.Avatar,
	}
	receiverRequestUser := requestmdl.RequestUser{
		Id:     user,
		Name:   receiver.Name,
		Avatar: receiver.Avatar,
	}
	groupImageUrl := ""
	if group.ImageURL != nil {
		groupImageUrl = *group.ImageURL
	}
	groupRequest := requestmdl.RequestGroup{
		Id:       *group.Id,
		Name:     *group.Name,
		ImageUrl: groupImageUrl,
	}
	req := requestmdl.Request{
		Sender:   senderRequestUser,
		Receiver: receiverRequestUser,
		Group:    &groupRequest,
	}
	req.Process()
	err = biz.groupRepo.CreateRequest(ctx, &req)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not create request"))
	}

	go func() {

		groupName := ""
		if group.Name != nil {
			groupName = *group.Name
		}

		e := biz.notification.CreateReceiveGroupRequestNotification(context.Background(), user,
			&notimodel.NotificationObject{
				Id:    user,
				Name:  receiver.Name,
				Image: &receiver.Avatar,
				Type:  notimodel.User,
			},
			&notimodel.NotificationObject{
				Id:    *req.Id,
				Name:  "",
				Image: nil,
				Type:  notimodel.Request,
			},
			&notimodel.NotificationObject{
				Id:    *group.Id,
				Name:  groupName,
				Image: group.ImageURL,
				Type:  notimodel.Group,
			},
			&notimodel.NotificationObject{
				Id:    requesterId,
				Name:  requester.Name,
				Image: &requester.Avatar,
				Type:  notimodel.User,
			})
		if e != nil {
			log.Error().Stack().Err(e).Msg("create receive group notification failed")
		}
	}()

	return nil
}
