package usermodel

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"time"
)

type User struct {
	common.MongoModel `json:"inline" bson:"inline"`
	Name              string     `json:"name" bson:"name"`
	Email             string     `json:"email" bson:"email"`
	Password          string     `bson:"password" json:"-"`
	Avatar            string     `json:"avatar" bson:"avatar"`
	Address           string     `bson:"address" json:"address"`
	Phone             string     `json:"phone" bson:"phone"`
	Gender            string     `json:"gender" bson:"gender"`
	Birthday          *time.Time `json:"birthday" bson:"birthday"`
}

func (User) EntityName() string {
	return "User"
}

func (User) CollectionName() string {
	return "users"
}

func (u *User) Process() {
}

func (u *User) GetId() string {
	return *u.Id
}
