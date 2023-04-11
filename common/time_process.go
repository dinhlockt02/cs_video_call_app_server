package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func MongoProcessTime(goTime *time.Time, mongoTime *primitive.DateTime) (*time.Time, *primitive.DateTime) {
	var t *time.Time
	if mongoTime == nil && goTime != nil {
		t = goTime
	} else if mongoTime != nil && goTime == nil {
		tempT := mongoTime.Time()
		t = &tempT
	} else {
		t = nil
	}
	var mt primitive.DateTime
	if t != nil {
		mt = primitive.NewDateTimeFromTime(*t)
	}
	return t, &mt
}
