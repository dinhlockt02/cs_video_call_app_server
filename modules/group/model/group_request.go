package groupmdl

import "github.com/dinhlockt02/cs_video_call_app_server/common"

type Filter string

const (
	Sent    Filter = "sent"
	Receive        = "receive"
)

type GroupRequest struct {
	*common.MongoId `json:",inline,omitempty" bson:",inline,omitempty"`
	Name            string   `bson:"name" json:"name"`
	Members         []string `bson:"members" json:"members"`
	ImageUrl        string   `json:"image_url" bson:"image_url"`
}

func (GroupRequest) CollectionName() string {
	return common.GroupCollectionName
}
