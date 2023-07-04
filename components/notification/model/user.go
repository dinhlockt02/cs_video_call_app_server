package notimodel

import "github.com/dinhlockt02/cs_video_call_app_server/common"

type NotificationUser struct {
	Notification bool `bson:"notification" json:"notification"`
}

func (NotificationUser) CollectionName() string {
	return common.UserCollectionName
}
