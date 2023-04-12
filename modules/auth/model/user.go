package authmodel

import (
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
)

type User struct {
	Id             string `json:"id" bson:"_id"`
	Email          string `json:"email" bson:"email"`
	Password       string `json:"password" json:"password"`
	EmailVerified  bool   `json:"email_verified" bson:"email_verified"`
	ProfileUpdated bool   `json:"profile_updated" bson:"profile_updated"`
}

func (User) CollectionName() string {
	return "users"
}

var ErrUserNotFound = common.ErrEntityNotFound("User", errors.New("user not found"))
