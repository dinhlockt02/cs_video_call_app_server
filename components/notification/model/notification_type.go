package notimodel

type NotificationObjectType string

const (
	User NotificationObjectType = "user"
)

type NotificationActionType string

const (
	AcceptRequest        NotificationActionType = "accept-request"
	ReceiveFriendRequest                        = "send-friend-request"
)
