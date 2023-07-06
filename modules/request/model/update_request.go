package requestmdl

import "github.com/dinhlockt02/cs_video_call_app_server/common"

type UpdateRequest struct {
	Sender   *RequestUser  `json:"sender,omitempty" bson:"sender,omitempty"`
	Receiver *RequestUser  `json:"receiver,omitempty" bson:"receiver,omitempty"`
	Group    *RequestGroup `bson:"group,omitempty" json:"group,omitempty"`
}

func (UpdateRequest) CollectionName() string {
	return common.RequestsCollectionName
}
