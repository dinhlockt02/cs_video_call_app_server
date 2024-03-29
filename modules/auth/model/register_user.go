package authmodel

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type RegisterUser struct {
	common.MongoId        `bson:",inline"`
	common.MongoCreatedAt `bson:",inline" json:",inline"`
	common.MongoUpdatedAt `bson:",inline" json:",inline"`
	Email                 string `json:"email" bson:"email"`
	Password              string `json:"password" bson:"password"`
	EmailVerified         bool   `json:"email_verified" bson:"email_verified"`
	ProfileUpdated        bool   `json:"profile_updated" bson:"profile_updated"`
}

func (RegisterUser) CollectionName() string {
	return common.UserCollectionName
}

func (u *RegisterUser) Process() error {
	var errs = make([]error, 0)

	if !common.EmailRegexp.Match([]byte(u.Email)) {
		errs = append(errs, errors.New("invalid email"))
	}

	if len(strings.TrimSpace(u.Password)) < 8 {
		errs = append(errs, errors.New("password must be at least 8 character"))
	}

	if len(strings.TrimSpace(u.Password)) > 50 {
		errs = append(errs, errors.New("password must be at most 50 character"))
	}

	now := time.Now()
	u.CreatedAt = &now
	u.UpdatedAt = &now

	if len(errs) > 0 {
		err := errs[0]
		for i := 1; i < len(errs); i++ {
			err = errors.Wrap(err, errs[i].Error())
		}
		return errors.Wrap(err, "validation error")
	}
	return nil
}
