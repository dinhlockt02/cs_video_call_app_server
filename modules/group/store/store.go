package groupstore

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	Create(ctx context.Context, group *groupmdl.Group) error
	List(
		ctx context.Context,
		filter map[string]interface{},
	) ([]groupmdl.Group, error)

	FindUser(
		ctx context.Context,
		filter map[string]interface{},
	) (*groupmdl.User, error)
	UpdateUser(
		ctx context.Context,
		filter map[string]interface{},
		updatedUser *groupmdl.User,
	) error
	FindGroup(
		ctx context.Context,
		filter map[string]interface{},
	) (*groupmdl.Group, error)
	UpdateGroup(
		ctx context.Context,
		filter map[string]interface{},
		updatedGroup *groupmdl.Group,
	) error
	FindUsers(
		ctx context.Context,
		filter map[string]interface{},
	) ([]groupmdl.User, error)
	DeleteOne(
		ctx context.Context,
		filter map[string]interface{},
	) error
}

type mongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *mongoStore {
	return &mongoStore{database: database}
}
