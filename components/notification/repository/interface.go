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

	// CreateReceiveFriendRequestNotification is a method that will create, store and push notification
	//
	// It should be used when the Subject (aka owner) received the friend request from Prep's
	CreateReceiveFriendRequestNotification(
		ctx context.Context,
		owner string,
		subject *notimodel.NotificationObject,
		prep *notimodel.NotificationObject,
	) error

	// CreateIncomingCallNotification should be used when the Subject call the Direct (aka owner) in a room (Prep)
	CreateIncomingCallNotification(
		ctx context.Context,
		owner string,
		subject *notimodel.NotificationObject,
		direct *notimodel.NotificationObject,
		prep *notimodel.NotificationObject,
	) error

	// CreateRejectIncomingCallNotification should be used when the Subject reject the Direct (aka owner) in a room (Prep)
	CreateRejectIncomingCallNotification(
		ctx context.Context,
		owner string,
		subject *notimodel.NotificationObject,
		direct *notimodel.NotificationObject,
		prep *notimodel.NotificationObject,
	) error

	// CreateAbandonIncomingCallNotification should be used when the Subject abandon call before the Direct (aka owner) answer in a room (Prep)
	CreateAbandonIncomingCallNotification(
		ctx context.Context,
		owner string,
		subject *notimodel.NotificationObject,
		direct *notimodel.NotificationObject,
		prep *notimodel.NotificationObject,
	) error
}
