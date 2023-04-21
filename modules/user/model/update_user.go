package usermodel

import (
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"strings"
	"time"
)

type UpdateUser struct {
	common.MongoModel              `json:"inline" bson:"inline"`
	common.MongoUpdatedAtTimestamp `json:"inline" bson:"inline"`
	Name                           string     `json:"name" bson:"name"`
	Avatar                         string     `json:"avatar" bson:"avatar"`
	Address                        string     `bson:"address" json:"address"`
	Phone                          string     `json:"phone" bson:"phone"`
	Gender                         string     `json:"gender" bson:"gender"`
	Birthday                       *time.Time `json:"birthday" bson:"birthday"`
	//MongoBirthday     *primitive.DateTime `bson:"birthday" json:"-"`
}

func (UpdateUser) EntityName() string {
	return "User"
}

func (UpdateUser) CollectionName() string {
	return "users"
}

func (u *UpdateUser) Process() error {
	var errs = make([]error, 0)
	if strings.TrimSpace(u.Name) == "" {
		errs = append(errs, errors.New("name must not be empty"))
	}

	if !common.URLRegexp.Match([]byte(u.Avatar)) {
		errs = append(errs, errors.New("invalid image url"))
	}

	if gender := strings.TrimSpace(string(u.Gender)); gender == "" {
		errs = append(errs, errors.New("gender must not be empty"))
	} else if gender != common.Male && gender != common.Female {
		errs = append(errs, errors.New("gender must be male or female"))
	}

	if strings.TrimSpace(u.Address) == "" {
		errs = append(errs, errors.New("address must not be empty"))
	}

	if strings.TrimSpace(u.Phone) == "" {
		errs = append(errs, errors.New("phone number must not be empty"))
	}

	//if u.Birthday == nil {
	//	errs = append(errs, errors.New("birthday must not be empty"))
	//}

	now := time.Now()
	u.UpdatedAt = &now

	if len(errs) > 0 {
		return common.ValidationError(errs)
	}
	return nil
}
