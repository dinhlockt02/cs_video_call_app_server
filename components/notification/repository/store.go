package notirepo

import (
	"context"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
)

type NotificationStore interface {
	Create(ctx context.Context, data *notimodel.Notification) error
	FindDevice(ctx context.Context, filter map[string]interface{}) ([]notimodel.Device, error)
}
