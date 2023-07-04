package notirepo

import (
	"context"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notistore "github.com/dinhlockt02/cs_video_call_app_server/components/notification/store"
)

type INotificationRepository interface {
	List(ctx context.Context, filter map[string]interface{}) ([]notimodel.Notification, error)
	Delete(ctx context.Context, filter map[string]interface{}) error
	GetNotificationUser(ctx context.Context, userFilter map[string]interface{}) (*notimodel.NotificationUser, error)
	UpdateNotificationUser(ctx context.Context, userFilter map[string]interface{}, data *notimodel.NotificationUser) error
}

type NotificationRepository struct {
	store notistore.NotificationStore
}

func NewNotificationRepository(store notistore.NotificationStore) *NotificationRepository {
	return &NotificationRepository{
		store: store,
	}
}

func (n *NotificationRepository) List(ctx context.Context,
	filter map[string]interface{}) ([]notimodel.Notification, error) {
	return n.store.List(ctx, filter)
}

func (n *NotificationRepository) Delete(ctx context.Context, filter map[string]interface{}) error {
	return n.store.Delete(ctx, filter)
}

func (n *NotificationRepository) GetNotificationUser(ctx context.Context, userFilter map[string]interface{}) (*notimodel.NotificationUser, error) {
	return n.store.FindUser(ctx, userFilter)
}

func (n *NotificationRepository) UpdateNotificationUser(ctx context.Context, userFilter map[string]interface{}, data *notimodel.NotificationUser) error {
	return n.store.UpdateUser(ctx, userFilter, data)
}
