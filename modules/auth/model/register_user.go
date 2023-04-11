package authmodel

import (
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"strings"
	"time"
)

type RegisterUser struct {
	common.MongoModel `bson:"inline"`
	Email             string `json:"email" bson:"email"`
	Password          string `json:"password" bson:"password"`
}

func (RegisterUser) CollectionName() string {
	return "users"
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

	u.MongoTimestamp.Process()

	if len(errs) > 0 {
		return common.ValidationError(errs)
	}
	return nil
}
