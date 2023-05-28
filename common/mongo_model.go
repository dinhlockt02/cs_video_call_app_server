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

func GetOrFilter(filters ...map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"$or": filters,
	}
}

func GetAndFilter(filters ...map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"$and": filters,
	}
}

func GetExistsFilter(fieldName string, exists bool) map[string]interface{} {
	return map[string]interface{}{
		fieldName: map[string]interface{}{
			"$exists": exists,
		},
	}
}

func GetInFilter(fieldName string, values ...interface{}) map[string]interface{} {
	return map[string]interface{}{
		fieldName: map[string]interface{}{
			"$in": values,
		},
	}
}
