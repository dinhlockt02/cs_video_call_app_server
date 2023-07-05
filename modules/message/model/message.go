package messagemdl

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
)

type Message struct {
	common.MongoId        `bson:",inline" json:",inline"`
	common.MongoCreatedAt `bson:",inline" json:",inline"`
	Sender                *User  `bson:"sender" json:"sender"`
	Content               string `bson:"content" json:"content"`
}

func (*Message) GetCollectionName() string {
	return common.MessageCollectionName
}
