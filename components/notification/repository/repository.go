package notirepo

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	"github.com/rs/zerolog/log"
)

type NotificationRepository struct {
	service NotificationService
	store   NotificationStore
}

func NewNotificationRepository(service NotificationService, store NotificationStore) *NotificationRepository {
	return &NotificationRepository{
		service: service,
		store:   store,
	}
}

func (repo *NotificationRepository) CreateAcceptFriendNotification(
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

func (repo *NotificationRepository) CreateReceiveFriendRequestNotification(
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

func (repo *NotificationRepository) CreateIncomingCallNotification(
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

func (repo *NotificationRepository) CreateRejectIncomingCallNotification(ctx context.Context,
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

func (repo *NotificationRepository) CreateAbandonIncomingCallNotification(
	ctx context.Context,
	owner string,
	subject *notimodel.NotificationObject,
	direct *notimodel.NotificationObject,
	prep *notimodel.NotificationObject,
) error {
	noti := notimodel.
		NewNotificationBuilder(notimodel.RejectCall, owner).
		SetSubject(subject).
		SetPrep(prep).
		SetDirect(direct).
		Build()

	return repo.createNotification(ctx, noti)
}

func (repo *NotificationRepository) createNotification(ctx context.Context, noti *notimodel.Notification) error {
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
