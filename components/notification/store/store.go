package notistore

import (
	"context"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationStore interface {
	Create(ctx context.Context, data *notimodel.Notification) error
	FindDevice(ctx context.Context, filter map[string]interface{}) ([]notimodel.Device, error)
	List(ctx context.Context, filter map[string]interface{}) ([]notimodel.Notification, error)
	Delete(ctx context.Context, filter map[string]interface{}) error
	FindUser(ctx context.Context, filter map[string]interface{}) (*notimodel.NotificationUser, error)
	UpdateUser(ctx context.Context, filter map[string]interface{}, updatedData *notimodel.NotificationUser) error
}

type MongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *MongoStore {
	return &MongoStore{database: database}
}
