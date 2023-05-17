package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoId struct {
	Id *string `json:"id" bson:"_id,omitempty"`
}

type MongoCreatedAt struct {
	CreatedAt *time.Time `bson:"created_at" json:"created_at,omitempty"`
}

type MongoUpdatedAt struct {
	UpdatedAt *time.Time `bson:"updated_at" json:"update_at,omitempty"`
}

func ToObjectId(hex string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(hex)
}
