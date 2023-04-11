package usermodel

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	common.MongoModel `json:"inline" bson:"inline"`
	Name              string              `json:"name" bson:"name"`
	Email             string              `json:"email" bson:"email"`
	Password          string              `bson:"password" json:"-"`
	ImageURL          string              `json:"imageUrl" bson:"imageURL"`
	Address           string              `bson:"address" json:"address"`
	Phone             string              `json:"phone" bson:"phone"`
	Gender            string              `json:"gender" bson:"gender"`
	Birthday          *time.Time          `json:"birthday" bson:"-"`
	MongoBirthday     *primitive.DateTime `bson:"birthday" json:"-"`
}

func (User) EntityName() string {
	return "User"
}

func (User) CollectionName() string {
	return "users"
}

func (u *User) Process() {
	u.MongoModel.Process()
	u.Birthday, u.MongoBirthday = common.MongoProcessTime(u.Birthday, u.MongoBirthday)
}
