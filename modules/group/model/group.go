package groupmdl

import (
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
)

type Group struct {
	common.MongoId `json:",inline,omitempty" bson:",inline,omitempty"`
	Name           *string  `bson:"name,omitempty" json:"name,omitempty"`
	Members        []string `bson:"members,omitempty" json:"members,omitempty"`
	ImageURL       *string  `json:"image_url,omitempty" bson:"image_url,omitempty"`
}

func (Group) CollectionName() string {
	return common.GroupCollectionName
}

func (g *Group) Process() error {
	errs := common.ValidationError{}
	if g.ImageURL != nil && !common.URLRegexp.Match([]byte(*g.ImageURL)) {
		errs = append(errs, errors.New("invalid group image url"))
	}
	if len(*g.Name) == 0 {
		errs = append(errs, errors.New("invalid group name"))
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}
