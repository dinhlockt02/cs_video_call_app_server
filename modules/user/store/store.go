package userstore

import (
	"context"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error)
}

type mongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *mongoStore {
	return &mongoStore{database: database}
}
