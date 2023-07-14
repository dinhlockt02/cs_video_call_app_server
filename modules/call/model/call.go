package callmdl

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"time"
)

type Status string

const (
	Ended   Status = "ended"
	Missed  Status = "missed"
	Reject  Status = "reject"
	OnGoing Status = "on-going"
)

type User struct {
	Id     string `bson:"id" json:"id"`
	Name   string `bson:"name" json:"name"`
	Avatar string `json:"avatar" bson:"avatar"`
}

type Call struct {
	common.MongoId `json:",inline" bson:",inline"`
	Owner          []string   `json:"-" bson:"owner"`
	Caller         *User      `bson:"caller" json:"caller"`
	Callee         *User      `bson:"callee" json:"callee"`
	Status         Status     `bson:"status" json:"status"`
	CalledAt       *time.Time `bson:"called_at" json:"called_at"`
}

func (Call) CollectionName() string {
	return common.CallCollectionName
}
