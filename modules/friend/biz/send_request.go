package friendbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	friendrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/repository"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type SendRequestBiz struct {
	friendRepo   friendrepo.Repository
	notification notirepo.INotificationService
}

func NewSendRequestBiz(
	friendRepo friendrepo.Repository,
	notification notirepo.INotificationService,
) *SendRequestBiz {
	return &SendRequestBiz{
		friendRepo:   friendRepo,
		notification: notification,
	}
}

func (biz *SendRequestBiz) SendRequest(ctx context.Context, senderId string, receiverId string) error {
	log.Debug().Str("senderId", senderId).Str("receiverId", receiverId).Msg("get sent request")
	// Find exists request
	existedRequest, err := biz.friendRepo.FindRequest(ctx, senderId, receiverId)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find request"))
	}
	if existedRequest != nil {
		return common.ErrInvalidRequest(errors.New(friendmodel.RequestExists))
	}

	// Find sender
	filter, err := common.GetIdFilter(senderId)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid sender id"))
	}
	sender, err := biz.friendRepo.FindUser(ctx, filter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find sender"))
	}
	if sender == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.New(friendmodel.SenderNotFound))
	}

	for _, friend := range sender.Friends {
		if friend == receiverId {
			return common.ErrInvalidRequest(errors.New(friendmodel.HasBeenFriend))
		}
	}

	for _, blockedUser := range sender.BlockedUser {
		if blockedUser == receiverId {
			return common.ErrInvalidRequest(errors.New(friendmodel.UserBeBlocked))
		}
	}

	// Find Receiver
	filter, err = common.GetIdFilter(receiverId)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid receiver id"))
	}
	receiver, err := biz.friendRepo.FindUser(ctx, filter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find receiver"))
	}
	if receiver == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.New(friendmodel.ReceiverNotFound))
	}

	for _, blockedUser := range receiver.BlockedUser {
		if blockedUser == senderId {
			return common.ErrInvalidRequest(errors.New(friendmodel.UserBeBlocked))
		}
	}

	senderRequestUser := requestmdl.RequestUser{
		Id:     senderId,
		Name:   sender.Name,
		Avatar: sender.Avatar,
	}
	receiverRequestUser := requestmdl.RequestUser{
		Id:     receiverId,
		Name:   receiver.Name,
		Avatar: receiver.Avatar,
	}
	request := requestmdl.Request{
		Sender:   senderRequestUser,
		Receiver: receiverRequestUser,
	}
	request.Process()
	err = biz.friendRepo.CreateRequest(ctx, &request)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not create request"))
	}

	go func() {
		e := biz.notification.CreateReceiveFriendRequestNotification(context.Background(), receiverId,
			&notimodel.NotificationObject{
				Id:    receiverId,
				Name:  receiver.Name,
				Image: &receiver.Avatar,
				Type:  notimodel.User,
			}, &notimodel.NotificationObject{
				Id:    *request.Id,
				Name:  "",
				Image: new(string),
				Type:  notimodel.Request,
			},
			&notimodel.NotificationObject{
				Id:    senderId,
				Name:  sender.Name,
				Image: &sender.Avatar,
				Type:  notimodel.User,
			})
		if e != nil {
			log.Error().Err(e).Msg("send friend request notification failed")
		}
	}()

	return nil
}
