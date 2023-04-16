package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoModel struct {
	Id             *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	MongoTimestamp `bson:"inline"`
}

type MongoTimestamp struct {
	MongoCreatedAt *primitive.DateTime `json:"-" bson:"created_at,omitempty"`
	MongoUpdatedAt *primitive.DateTime `json:"-" bson:"updated_at,omitempty"`
	CreatedAt      *time.Time          `bson:"-" json:"created_at"`
	UpdatedAt      *time.Time          `bson:"-" json:"update_at"`
}

func (m *MongoTimestamp) Process() {
	m.CreatedAt, m.MongoCreatedAt = MongoProcessTime(m.CreatedAt, m.MongoCreatedAt)
	m.UpdatedAt, m.MongoUpdatedAt = MongoProcessTime(m.UpdatedAt, m.MongoUpdatedAt)
}
