package authmodel

import (
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	"strings"
)

type LoginUser struct {
	Email    string              `json:"email" bson:"email"`
	Password string              `json:"password" json:"password"`
	Device   *devicemodel.Device `json:"device" bson:"-"`
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

	if u.Device == nil {
		errs = append(errs, errors.New("device must not be null"))
	}

	if err := u.Device.Process(); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return common.ValidationError(errs)
	}
	return nil
}
