package messagestore

import (
	"context"
	messagemdl "github.com/dinhlockt02/cs_video_call_app_server/modules/message/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	Create(ctx context.Context, message *messagemdl.Message) error
	DeleteOne(ctx context.Context, filter map[string]interface{}) error
	List(ctx context.Context, filter map[string]interface{}) ([]messagemdl.Message, error)
	Get(ctx context.Context, filter map[string]interface{}) (*messagemdl.Message, error)
}

type MongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *MongoStore {
	return &MongoStore{database: database}
}
