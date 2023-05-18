package notimodel

import (
	"fmt"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"time"
)

type Notification struct {
	common.MongoId        `bson:",inline" json:",inline"`
	common.MongoCreatedAt `json:",inline" bson:",inline"`

	// Owner is a string that represent the id of the user who receive notification
	Owner string `bson:"owner" json:"owner"`
	// Subject is a NotificationObject that represent the object do the Action
	Subject *NotificationObject `bson:"subject" json:"subject"`
	// Direct is a NotificationObject that represent the object was directly affected by the Action
	Direct *NotificationObject `bson:"direct" json:"direct"`
	// Indirect is a NotificationObject that represent the object was indirectly affected by the Action
	Indirect *NotificationObject `bson:"indirect" json:"indirect"`
	// Indirect is a NotificationObject that represent the object was appear in the action with a prep (in, for, of)
	Prep *NotificationObject `bson:"prep" json:"prep"`
	// Action is a string has type of NotificationActionType
	Action NotificationActionType `json:"action" bson:"action"`
}

func (Notification) CollectionName() string {
	return "notifications"
}

type notificationBuilder struct {
	Owner    string
	Subject  *NotificationObject
	Direct   *NotificationObject
	Indirect *NotificationObject
	Prep     *NotificationObject
	Action   NotificationActionType
}

func NewNotificationBuilder(action NotificationActionType, owner string) *notificationBuilder {
	result := new(notificationBuilder)
	result.Action = action
	result.Owner = owner
	return result
}

func (builder *notificationBuilder) SetSubject(object *NotificationObject) *notificationBuilder {
	builder.Subject = object
	return builder
}

func (builder *notificationBuilder) SetDirect(object *NotificationObject) *notificationBuilder {
	builder.Direct = object
	return builder
}

func (builder *notificationBuilder) SetIndirect(object *NotificationObject) *notificationBuilder {
	builder.Indirect = object
	return builder
}

func (builder *notificationBuilder) SetPrep(object *NotificationObject) *notificationBuilder {
	builder.Prep = object
	return builder
}

func (builder *notificationBuilder) Build() *Notification {
	now := time.Now()
	return &Notification{
		MongoCreatedAt: common.MongoCreatedAt{CreatedAt: &now},
		Subject:        builder.Subject,
		Direct:         builder.Direct,
		Indirect:       builder.Indirect,
		Prep:           builder.Prep,
		Owner:          builder.Owner,
		Action:         builder.Action,
	}
}

// GetMessage is a function that will return 2 values respectively
// is title and the content of the notification.
func (n *Notification) GetMessage() (title string, body string) {
	if n.Action == AcceptRequest {

	}
	switch n.Action {
	case AcceptRequest:
		return "Accept friend request", fmt.Sprintf("%s accept your friend request", n.Subject.Name)
	case ReceiveFriendRequest:
		return "Friend request received", fmt.Sprintf("%s want to be friend with you", n.Prep.Name)
	default:
		return "", ""
	}

}
