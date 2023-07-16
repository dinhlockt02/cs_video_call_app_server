package friendmodel

import (
	"time"

	"github.com/dinhlockt02/cs_video_call_app_server/common"
)

type FriendUser struct {
	Id         *string    `json:"id" bson:"_id,omitempty"`
	Email      string     `json:"email" bson:"email"`
	Avatar     string     `json:"avatar" bson:"avatar"`
	Name       string     `json:"name" bson:"name"`
	LastSeenAt *time.Time `json:"last_seen_at" bson:"last_seen_at"`
}

func (FriendUser) CollectionName() string {
	return common.UserCollectionName
}
