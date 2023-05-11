package friendmodel

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"time"
)

type RequestUser struct {
	Id     string `bson:"id" json:"id"`
	Name   string `bson:"name" json:"name"`
	Avatar string `json:"avatar" bson:"avatar"`
}

type Request struct {
	Id                             *string     `json:"-" bson:"_id,omitempty"`
	Sender                         RequestUser `json:"sender" bson:"sender"`
	Receiver                       RequestUser `json:"receiver" bson:"receiver"`
	common.MongoCreatedAtTimestamp `json:",inline" bson:",inline"`
}

func (Request) CollectionName() string {
	return "requests"
}
func (r *Request) Process() {
	now := time.Now()
	r.CreatedAt = &now
}
