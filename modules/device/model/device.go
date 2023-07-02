package devicemodel

import (
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"strings"
	"time"
)

type Device struct {
	common.MongoId        `bson:",inline"`
	common.MongoCreatedAt `bson:",inline" json:",inline"`
	common.MongoUpdatedAt `bson:",inline" json:",inline"`
	Name                  string `bson:"name" json:"name"`
	UserId                string `bson:"user_id,omitempty" json:"-"`
	PushNotificationToken string `json:"push_notification_token" bson:"push_notification_token"`
}

func (d *Device) Process() error {
	var errs = make([]error, 0)

	if len(strings.TrimSpace(d.Name)) == 0 {
		errs = append(errs, errors.New("device name must not be empty"))
	}

	if len(errs) > 0 {
		return common.ValidationError(errs)
	}

	now := time.Now()
	d.CreatedAt = &now
	d.UpdatedAt = &now

	return nil
}

func (Device) CollectionName() string {
	return common.DevicesCollectionName
}
