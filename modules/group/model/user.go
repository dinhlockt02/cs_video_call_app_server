package groupmdl

import "github.com/dinhlockt02/cs_video_call_app_server/common"

type User struct {
	common.MongoId `json:",inline" bson:",inline"`
	Groups         []string `bson:"groups"`
	Avatar         string   `bson:"avatar" json:"avatar"`
	Name           string   `json:"name" bson:"name"`
}

func (User) CollectionName() string {
	return common.UserCollectionName
}
