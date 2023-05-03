package authmiddleware

import "github.com/dinhlockt02/cs_video_call_app_server/common"

type Device struct {
	common.MongoModel `bson:",inline" json:",inline"`
	UserId            string `json:"user_id" bson:"user_id"`
}

func (Device) CollectionName() string {
	return "devices"
}

func (u *Device) GetId() string {
	return u.UserId
}

func (u *Device) GetDeviceId() string {
	return *u.Id
}
