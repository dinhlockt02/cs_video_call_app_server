package notirepo

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	"github.com/rs/zerolog/log"
)

type notificationRepository struct {
	service NotificationService
	store   NotificationStore
}

func NewNotificationRepository(service NotificationService, store NotificationStore) *notificationRepository {
	return &notificationRepository{
		service: service,
		store:   store,
	}
}

func (repo *notificationRepository) CreateAcceptFriendNotification(
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

func (repo *notificationRepository) CreateReceiveFriendRequestNotification(
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

func (repo *notificationRepository) createNotification(ctx context.Context, noti *notimodel.Notification) error {
	err := repo.store.Create(ctx, noti)
	if err != nil {
		return common.ErrInternal(err)
	}

	go func() {
		devices, e := repo.store.FindDevice(ctx, map[string]interface{}{
			"user_id": noti.Owner,
		})
		if e != nil {
			log.Err(e)
		}

		tokens := make([]string, len(devices))

		for i, _ := range devices {
			tokens[i] = devices[i].PushNotificationToken
		}

		e = repo.service.Push(context.Background(), tokens, noti)
		if e != nil {
			log.Err(e)
		}
	}()

	return nil
}
