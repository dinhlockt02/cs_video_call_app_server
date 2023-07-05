package messagemdl

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"time"
)

type Message struct {
	common.MongoId        `bson:",inline" json:",inline"`
	common.MongoCreatedAt `bson:",inline" json:",inline"`
	GroupId               string     `json:"group_id" bson:"group_id"`
	SenderId              *string    `json:"sender_id,omitempty"`
	Sender                *User      `bson:"sender" json:"sender"`
	Content               string     `bson:"content" json:"content"`
	SentAt                *time.Time `bson:"sent_at" json:"sent_at"`
}

func (*Message) CollectionName() string {
	return common.MessageCollectionName
}

func (m *Message) Process() {
	m.CreatedAt = common.Ptr(time.Now())
}
