package notirepo

import (
	"context"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
)

type NotificationRepository interface {
	// CreateAcceptFriendNotification is a method that will create, store and push notification
	//
	// It should be used when the subject accept the indirect (aka owner)'s friend request
	CreateAcceptFriendNotification(
		ctx context.Context,
		owner string,
		subject *notimodel.NotificationObject,
		indirect *notimodel.NotificationObject,
	) error
}
