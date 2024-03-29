package friendmodel

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"time"
)

type Relation string

const (
	Self     Relation = "self"
	Friend   Relation = "friend"
	Received Relation = "received"
	Sent     Relation = "sent"
	Blocked  Relation = "blocked"
	Non      Relation = "non"
)

type User struct {
	Id          *string    `json:"id" bson:"_id,omitempty"`
	UpdatedAt   *time.Time `bson:"updated_at" json:"update_at,omitempty"`
	Friends     []string   `json:"friends" bson:"friends"`
	Groups      []string   `json:"-" bson:"groups"`
	BlockedUser []string   `bson:"blocked_user"`
	Avatar      string     `json:"avatar" bson:"avatar"`
	Name        string     `json:"name" bson:"name"`
	Relation    Relation   `json:"relation"`
}

func (*User) CollectionName() string {
	return common.UserCollectionName
}

func (u *User) Process() {
	now := time.Now()
	u.UpdatedAt = &now
}
