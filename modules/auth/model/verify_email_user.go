package authmodel

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"time"
)

type EmailVerifiedUser struct {
	EmailVerified                  bool `json:"email_verified" bson:"email_verified"`
	common.MongoUpdatedAtTimestamp `bson:",inline"`
}

func (EmailVerifiedUser) CollectionName() string {
	return "users"
}

func (u *EmailVerifiedUser) Process() {
	now := time.Now()
	u.UpdatedAt = &now
	u.EmailVerified = true
}
