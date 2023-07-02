package authmodel

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/pkg/errors"
	"strings"
)

type LoginUser struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

func (LoginUser) CollectionName() string {
	return "users"
}

func (u *LoginUser) Process() error {
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

	if len(errs) > 0 {
		err := errs[0]
		for i := 1; i < len(errs); i++ {
			err = errors.Wrap(err, errs[i].Error())
		}
		return errors.Wrap(err, "validation error")
	}
	return nil
}
