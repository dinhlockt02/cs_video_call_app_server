package authmodel

import (
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"time"
)

type RegisterFirebaseUser struct {
	common.MongoModel `bson:"inline"`
	Email             string `json:"email" bson:"email"`
}

func (RegisterFirebaseUser) CollectionName() string {
	return "users"
}

func (u *RegisterFirebaseUser) Process() error {
	var errs = make([]error, 0)

	if !common.EmailRegexp.Match([]byte(u.Email)) {
		errs = append(errs, errors.New("invalid email"))
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
