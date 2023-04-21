package devicemodel

import (
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"strings"
	"time"
)

type Device struct {
	common.MongoModel              `bson:"inline"`
	common.MongoCreatedAtTimestamp `bson:"inline" json:"inline"`
	common.MongoUpdatedAtTimestamp `bson:"inline" json:"inline"`
	Name                           string `bson:"name" json:"name"`
	UserId                         string `bson:"user_id" json:"-"`
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
	return "devices"
}
