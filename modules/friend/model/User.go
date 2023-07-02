package friendmodel

import (
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
	BlockedUser []string   `bson:"blocked_user"`
	Avatar      string     `json:"avatar" bson:"avatar"`
	Name        string     `json:"name" bson:"name"`
	Relation    Relation   `json:"relation"`
}

func (User) CollectionName() string {
	return "users"
}

func (u *User) Process() {
	now := time.Now()
	u.UpdatedAt = &now
}
