package common

import (
	"time"
)

type MongoModel struct {
	Id *string `json:"id" bson:"_id,omitempty"`
}

//type MongoTimestamp struct {
//	MongoCreatedAt *primitive.DateTime `json:"-" bson:"created_at,omitempty"`
//	MongoUpdatedAt *primitive.DateTime `json:"-" bson:"updated_at,omitempty"`
//	CreatedAt      *time.Time          `bson:"-" json:"created_at"`
//	UpdatedAt      *time.Time          `bson:"-" json:"update_at"`
//}

type MongoCreatedAtTimestamp struct {
	CreatedAt *time.Time `bson:"created_at" json:"created_at,omitempty"`
}

type MongoUpdatedAtTimestamp struct {
	UpdatedAt *time.Time `bson:"updated_at" json:"update_at,omitempty"`
}
