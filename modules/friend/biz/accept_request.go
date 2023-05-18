package friendbiz

import (
	"context"
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AcceptRequestFriendStore interface {
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
	UpdateUser(ctx context.Context, filter map[string]interface{}, updatedUser *friendmodel.User) error
	DeleteRequest(ctx context.Context, requestId string) error
	FindRequest(ctx context.Context, userId string, otherId string) (*friendmodel.Request, error)
}

type acceptRequestBiz struct {
	friendStore  AcceptRequestFriendStore
	notification notirepo.NotificationRepository
}

func NewAcceptRequestBiz(
	friendStore AcceptRequestFriendStore,
	notification notirepo.NotificationRepository,
) *acceptRequestBiz {
	return &acceptRequestBiz{
		friendStore:  friendStore,
		notification: notification,
	}
}

func (biz *acceptRequestBiz) AcceptRequest(ctx context.Context, senderId string, receiverId string) error {
	existedRequest, err := biz.friendStore.FindRequest(ctx, senderId, receiverId)
	if err != nil {
		return err
	}
	if existedRequest == nil {
		return common.ErrInvalidRequest(friendmodel.ErrRequestNotFound)
	}
	id, _ := primitive.ObjectIDFromHex(senderId)
	sender, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})

	if err != nil {
		return err
	}
	if sender == nil {
		return common.ErrEntityNotFound("User", errors.New("sender not found"))
	}
	id, _ = primitive.ObjectIDFromHex(receiverId)
	receiver, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})
	if receiver == nil {
		return common.ErrEntityNotFound("User", errors.New("receiver not found"))
	}
	sender.Friends = append(sender.Friends, receiverId)
	receiver.Friends = append(receiver.Friends, senderId)

	id, _ = primitive.ObjectIDFromHex(senderId)
	err = biz.friendStore.UpdateUser(ctx, map[string]interface{}{
		"_id": id,
	}, sender)
	if err != nil {
		return err
	}

	id, _ = primitive.ObjectIDFromHex(receiverId)
	err = biz.friendStore.UpdateUser(ctx, map[string]interface{}{
		"_id": id,
	}, receiver)
	if err != nil {
		return err
	}
	err = biz.friendStore.DeleteRequest(ctx, *existedRequest.Id)
	if err != nil {
		return err
	}

	go func() {
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
			log.Err(e)
		}
	}()
	return nil
}
