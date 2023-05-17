package notirepo

import (
	"context"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
)

type NotificationService interface {
	Push(ctx context.Context, token []string, notification *notimodel.Notification) error
}
