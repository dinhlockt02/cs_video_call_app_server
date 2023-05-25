package requeststore

import (
	"context"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	FindRequests(ctx context.Context, filter map[string]interface{}) ([]requestmdl.Request, error)
	FindRequest(ctx context.Context, filter map[string]interface{}) (*requestmdl.Request, error)
	DeleteRequest(ctx context.Context, filter map[string]interface{}) error
	CreateRequest(ctx context.Context, request *requestmdl.Request) error
}

type mongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *mongoStore {
	return &mongoStore{database: database}
}
