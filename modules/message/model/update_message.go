package messagemdl

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
)

type UpdateMessage struct {
	Sender *User `bson:"sender,omitempty" json:"sender,omitempty"`
}

func (*UpdateMessage) CollectionName() string {
	return common.MessageCollectionName
}

func (m *UpdateMessage) Process() {
}
