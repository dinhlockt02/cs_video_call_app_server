package usermodel

import "time"

type UserDetail struct {
	Id                *string    `json:"id" bson:"_id,omitempty"`
	Name              string     `json:"name" bson:"name"`
	Email             string     `json:"email" bson:"email"`
	Avatar            string     `json:"avatar" bson:"avatar"`
	Phone             string     `json:"phone" bson:"phone"`
	Gender            string     `json:"gender" bson:"gender"`
	Birthday          *time.Time `json:"birthday" bson:"birthday"`
	CommonFriend      []string   `json:"-"`
	CommonFriendCount *int       `json:"common_friend_count,omitempty"`
	IsFriend          *bool      `json:"is_friend,omitempty"`
}

func NewUserDetail(u *User) *UserDetail {
	return &UserDetail{
		Id:                u.Id,
		Name:              u.Name,
		Email:             u.Email,
		Avatar:            u.Avatar,
		Phone:             u.Phone,
		Gender:            u.Gender,
		Birthday:          u.Birthday,
		CommonFriend:      nil,
		CommonFriendCount: new(int),
		IsFriend:          new(bool),
	}
}
