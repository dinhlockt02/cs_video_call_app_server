package friendbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FriendStore interface {
	CreateRequest(ctx context.Context, request *friendmodel.Request) error
	FindRequest(ctx context.Context, userId string, otherId string) (*friendmodel.Request, error)
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
}

type sendRequestBiz struct {
	friendStore  FriendStore
	notification notirepo.NotificationRepository
}

func NewSendRequestBiz(
	friendStore FriendStore,
	notification notirepo.NotificationRepository,
) *sendRequestBiz {
	return &sendRequestBiz{
		friendStore:  friendStore,
		notification: notification,
	}
}

func (biz *sendRequestBiz) SendRequest(ctx context.Context, senderId string, receiverId string) error {
	existedRequest, err := biz.friendStore.FindRequest(ctx, senderId, receiverId)
	if err != nil {
		return err
	}
	if existedRequest != nil {
		return common.ErrInvalidRequest(friendmodel.ErrRequestExists)
	}
	id, _ := primitive.ObjectIDFromHex(senderId)
	sender, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})
	if err != nil {
		return err
	}
	if sender == nil {
		return common.ErrEntityNotFound("User", nil)
	}

	for i := range sender.Friends {
		if sender.Friends[i] == receiverId {
			return common.ErrInvalidRequest(friendmodel.ErrHasBeenFriend)
		}
	}

	id, _ = primitive.ObjectIDFromHex(receiverId)
	receiver, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})
	if err != nil {
		return err
	}
	if receiver == nil {
		return common.ErrEntityNotFound("User", nil)
	}
	senderRequestUser := friendmodel.RequestUser{
		Id:     senderId,
		Name:   sender.Name,
		Avatar: sender.Avatar,
	}
	receiverRequestUser := friendmodel.RequestUser{
		Id:     receiverId,
		Name:   receiver.Name,
		Avatar: receiver.Avatar,
	}
	request := friendmodel.Request{
		Sender:   senderRequestUser,
		Receiver: receiverRequestUser,
	}
	request.Process()
	err = biz.friendStore.CreateRequest(ctx, &request)
	if err != nil {
		return err
	}

	go func() {
		e := biz.notification.CreateReceiveFriendRequestNotification(context.Background(), receiverId, &notimodel.NotificationObject{
			Id:    receiverId,
			Name:  receiver.Name,
			Image: &receiver.Avatar,
			Type:  notimodel.User,
		}, &notimodel.NotificationObject{
			Id:    senderId,
			Name:  sender.Name,
			Image: &sender.Avatar,
			Type:  notimodel.User,
		})
		if e != nil {
			log.Err(e)
		}
	}()

	return nil
}
