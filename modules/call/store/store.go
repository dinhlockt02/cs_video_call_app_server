package callstore

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	Create(ctx context.Context, data *callmdl.Call) error
	FindCall(ctx context.Context, filter map[string]interface{}) (*callmdl.Call, error)
	Update(ctx context.Context, filter map[string]interface{}, data *callmdl.Call) error
}

type MongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *MongoStore {
	return &MongoStore{database: database}
}
