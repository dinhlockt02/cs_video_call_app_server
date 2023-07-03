package callmdl

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
)

type UpdateCall struct {
	Caller *User  `bson:"caller,omitempty" json:"caller"`
	Callee *User  `bson:"callee,omitempty" json:"callee"`
	Status Status `bson:"status,omitempty" json:"status"`
}

func (UpdateCall) CollectionName() string {
	return common.CallCollectionName
}
