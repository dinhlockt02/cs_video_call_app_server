package groupbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
	"github.com/pkg/errors"
)

type sendGroupRequestBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.NotificationRepository
}

func NewSendGroupRequestBiz(groupRepo grouprepo.Repository, notification notirepo.NotificationRepository) *sendGroupRequestBiz {
	return &sendGroupRequestBiz{groupRepo: groupRepo, notification: notification}
}

// SendRequest send a group invitation request to user.
func (biz *sendGroupRequestBiz) SendRequest(ctx context.Context, requesterId string, user string, group *groupmdl.Group) error {

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
	if group.ImageUrl != nil {
		groupImageUrl = *group.ImageUrl
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
		// TODO: Push notification group request
		//e := biz.notification.CreateReceiveFriendRequestNotification(
		//	context.Background(), receiverId, &notimodel.NotificationObject{
		//		Id:    receiverId,
		//		Name:  receiver.Name,
		//		Image: &receiver.Avatar,
		//		Type:  notimodel.User,
		//	}, &notimodel.NotificationObject{
		//		Id:    senderId,
		//		Name:  sender.Name,
		//		Image: &sender.Avatar,
		//		Type:  notimodel.User,
		//	})
		//if e != nil {
		//	log.Err(e)
		//}
	}()

	return nil
}
