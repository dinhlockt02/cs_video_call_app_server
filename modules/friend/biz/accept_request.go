package friendbiz

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	friendrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/repository"
)

type AcceptRequestBiz struct {
	friendRepository friendrepo.Repository
	notification     notirepo.NotificationRepository
}

func NewAcceptRequestBiz(
	friendRepository friendrepo.Repository,
	notification notirepo.NotificationRepository,
) *AcceptRequestBiz {
	return &AcceptRequestBiz{
		friendRepository: friendRepository,
		notification:     notification,
	}
}

func (biz *AcceptRequestBiz) AcceptRequest(ctx context.Context, senderId string, receiverId string) error {
	log.Debug().Str("senderId", senderId).Str("receiverId", receiverId).Msg("Accept friend request")
	// Find exists request
	existedRequest, err := biz.friendRepository.FindRequest(ctx, senderId, receiverId)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find request"))
	}
	if existedRequest == nil {
		return common.ErrInvalidRequest(errors.New(friendmodel.RequestNotFound))
	}

	// Find sender
	senderFilter, err := common.GetIdFilter(senderId)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid sender id"))
	}
	sender, err := biz.friendRepository.FindUser(ctx, senderFilter)

	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find sender"))
	}

	if sender == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.New(friendmodel.SenderNotFound))
	}

	// Find Receiver
	receiverFilter, err := common.GetIdFilter(receiverId)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid receiver id"))
	}
	receiver, err := biz.friendRepository.FindUser(ctx, receiverFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find receiver"))
	}
	if receiver == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.New(friendmodel.ReceiverNotFound))
	}

	// Update Sender
	sender.Friends = append(sender.Friends, receiverId)

	err = biz.friendRepository.UpdateUser(ctx, senderFilter, sender)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not update sender"))
	}

	// Update Receiver
	receiver.Friends = append(receiver.Friends, senderId)
	err = biz.friendRepository.UpdateUser(ctx, receiverFilter, receiver)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not update receiver"))
	}

	// Delete request
	requestFilter, err := common.GetIdFilter(*existedRequest.Id)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "invalid request filter"))
	}
	err = biz.friendRepository.DeleteRequest(ctx, requestFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not delete request"))
	}

	go func() {
		// TODO: change to pubsub model
		e := biz.notification.CreateAcceptFriendNotification(ctx, senderId, &notimodel.NotificationObject{
			Id:    senderId,
			Name:  sender.Name,
			Image: &sender.Avatar,
			Type:  notimodel.User,
		}, &notimodel.NotificationObject{
			Id:    receiverId,
			Name:  receiver.Name,
			Image: &receiver.Avatar,
			Type:  notimodel.User,
		})
		if e != nil {
			log.Error().Err(e).Msg("create accept friend notification failed")
		}
	}()
	return nil
}
