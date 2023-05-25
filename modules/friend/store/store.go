package friendstore

import (
	"context"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	FindFriend(ctx context.Context, filter map[string]interface{}) ([]friendmodel.FriendUser, error)
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
	UpdateUser(ctx context.Context, filter map[string]interface{}, updatedUser *friendmodel.User) error
}

type mongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *mongoStore {
	return &mongoStore{database: database}
}
