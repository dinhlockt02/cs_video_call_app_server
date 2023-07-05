package notirepo

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notiservice "github.com/dinhlockt02/cs_video_call_app_server/components/notification/service"
	notistore "github.com/dinhlockt02/cs_video_call_app_server/components/notification/store"
	"github.com/rs/zerolog/log"
)

type INotificationService interface {
	// CreateAcceptFriendNotification is a method that will create, store and push notification
	//
	// It should be used when the subject accept the indirect (aka owner)'s friend request
	CreateAcceptFriendNotification(
		ctx context.Context,
		owner string,
		subject *notimodel.NotificationObject,
		indirect *notimodel.NotificationObject,
	) error

	// CreateReceiveFriendRequestNotification is a method that will create, store and push notification
	//
	// It should be used when the Subject (aka owner) received the friend request from Prep's
	CreateReceiveFriendRequestNotification(
		ctx context.Context,
		owner string,
		subject *notimodel.NotificationObject,
		prep *notimodel.NotificationObject,
	) error

	// CreateReceiveGroupRequestNotification is a method that will create, store and push notification
	//
	// It should be used when
	// the Subject (aka owner) received the group request (Direct) to Group (Indirect) from Prep's
	CreateReceiveGroupRequestNotification(
		ctx context.Context,
		owner string,
		subject *notimodel.NotificationObject,
		direct *notimodel.NotificationObject,
		indirect *notimodel.NotificationObject,
		prep *notimodel.NotificationObject,
	) error

	// CreateIncomingCallNotification should be used when
	// the Subject call the Direct (aka owner) in a room (Prep)
	CreateIncomingCallNotification(
		ctx context.Context,
		owner string,
		subject *notimodel.NotificationObject,
		direct *notimodel.NotificationObject,
		prep *notimodel.NotificationObject,
	) error

	// CreateRejectIncomingCallNotification should be used when
	// the Subject reject the Direct (aka owner) in a room (Prep)
	CreateRejectIncomingCallNotification(
		ctx context.Context,
		owner string,
		subject *notimodel.NotificationObject,
		direct *notimodel.NotificationObject,
		prep *notimodel.NotificationObject,
	) error

	// CreateAbandonIncomingCallNotification should be used
	// when the Subject abandon call before the Direct (aka owner) answer in a room (Prep)
	CreateAbandonIncomingCallNotification(
		ctx context.Context,
		owner string,
		subject *notimodel.NotificationObject,
		direct *notimodel.NotificationObject,
		prep *notimodel.NotificationObject,
	) error
}

type NotificationService struct {
	service notiservice.NotificationService
	store   notistore.NotificationStore
}

func NewNotificationService(service notiservice.NotificationService,
	store notistore.NotificationStore) INotificationService {
	return &NotificationService{
		service: service,
		store:   store,
	}
}

func (repo *NotificationService) CreateAcceptFriendNotification(
	ctx context.Context,
	owner string,
	subject *notimodel.NotificationObject,
	indirect *notimodel.NotificationObject,
) error {
	noti := notimodel.
		NewNotificationBuilder(notimodel.AcceptRequest, owner).
		SetSubject(subject).
		SetIndirect(indirect).
		Build()

	return repo.createNotification(ctx, noti)
}

func (repo *NotificationService) CreateReceiveFriendRequestNotification(
	ctx context.Context,
	owner string,
	subject *notimodel.NotificationObject,
	prep *notimodel.NotificationObject,
) error {
	noti := notimodel.
		NewNotificationBuilder(notimodel.ReceiveFriendRequest, owner).
		SetSubject(subject).
		SetPrep(prep).
		Build()

	return repo.createNotification(ctx, noti)
}

func (repo *NotificationService) CreateIncomingCallNotification(
	ctx context.Context,
	owner string,
	subject *notimodel.NotificationObject,
	direct *notimodel.NotificationObject,
	prep *notimodel.NotificationObject,
) error {
	noti := notimodel.
		NewNotificationBuilder(notimodel.InComingCall, owner).
		SetSubject(subject).
		SetPrep(prep).
		SetDirect(direct).
		Build()

	return repo.createNotification(ctx, noti)
}

func (repo *NotificationService) CreateRejectIncomingCallNotification(ctx context.Context,
	owner string, subject *notimodel.NotificationObject,
	direct *notimodel.NotificationObject, prep *notimodel.NotificationObject) error {
	noti := notimodel.
		NewNotificationBuilder(notimodel.RejectCall, owner).
		SetSubject(subject).
		SetPrep(prep).
		SetDirect(direct).
		Build()

	return repo.createNotification(ctx, noti)
}

func (repo *NotificationService) CreateAbandonIncomingCallNotification(
	ctx context.Context,
	owner string,
	subject *notimodel.NotificationObject,
	direct *notimodel.NotificationObject,
	prep *notimodel.NotificationObject,
) error {
	noti := notimodel.
		NewNotificationBuilder(notimodel.AbandonCall, owner).
		SetSubject(subject).
		SetPrep(prep).
		SetDirect(direct).
		Build()

	return repo.createNotification(ctx, noti)
}

func (repo *NotificationService) CreateReceiveGroupRequestNotification(ctx context.Context, owner string, subject *notimodel.NotificationObject, direct *notimodel.NotificationObject, indirect *notimodel.NotificationObject, prep *notimodel.NotificationObject) error {
	noti := notimodel.
		NewNotificationBuilder(notimodel.ReceiveGroupRequest, owner).
		SetSubject(subject).
		SetDirect(direct).
		SetIndirect(indirect).
		SetPrep(prep).
		Build()

	return repo.createNotification(ctx, noti)
}

func (repo *NotificationService) createNotification(ctx context.Context, noti *notimodel.Notification) error {
	err := repo.store.Create(ctx, noti)
	if err != nil {
		return common.ErrInternal(err)
	}

	go func() {
		devices, e := repo.store.FindDevice(context.Background(), map[string]interface{}{
			"user_id": noti.Owner,
		})
		if e != nil {
			log.Err(e)
		}

		tokens := make([]string, len(devices))

		for i := range devices {
			tokens[i] = devices[i].PushNotificationToken
		}

		e = repo.service.Push(context.Background(), tokens, noti)
		if e != nil {
			log.Err(e)
		}
	}()

	return nil
}
