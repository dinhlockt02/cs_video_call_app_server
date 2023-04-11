package devicemodel

import (
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type Device struct {
	common.MongoModel `bson:"inline"`
	MongoExpiredAt    *primitive.DateTime `json:"-" bson:"expired_at"`
	ExpiredAt         *time.Time          `bson:"-" json:"expired_at"`
	Name              string              `bson:"name" json:"name"`
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

	d.MongoTimestamp.Process()

	d.ExpiredAt, d.MongoExpiredAt = common.MongoProcessTime(d.ExpiredAt, d.MongoExpiredAt)

	return nil
}

func (Device) CollectionName() string {
	return "devices"
}
