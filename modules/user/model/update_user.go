package usermodel

import (
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type UpdateUser struct {
	common.MongoModel `json:"inline" bson:"inline"`
	Name              string              `json:"name" bson:"name"`
	ImageUrl          string              `json:"imageUrl" bson:"imageUrl"`
	Address           string              `bson:"address" json:"address"`
	Phone             string              `json:"phone" bson:"phone"`
	Gender            string              `json:"gender" bson:"gender"`
	Birthday          *time.Time          `json:"birthday" bson:"-"`
	MongoBirthday     *primitive.DateTime `bson:"birthday" json:"-"`
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

	if !common.URLRegexp.Match([]byte(u.ImageUrl)) {
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

	u.MongoTimestamp.Process()
	u.Birthday, u.MongoBirthday = common.MongoProcessTime(u.Birthday, u.MongoBirthday)
	u.CreatedAt = nil
	u.MongoCreatedAt = nil

	if len(errs) > 0 {
		return common.ValidationError(errs)
	}
	return nil
}
