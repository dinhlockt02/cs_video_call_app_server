package notimodel

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
)

type Device struct {
	common.MongoId        `bson:",inline"`
	UserId                string `bson:"user_id" json:"-"`
	PushNotificationToken string `json:"push_notification_token" bson:"push_notification_token"`
}

func (Device) CollectionName() string {
	return "devices"
}
