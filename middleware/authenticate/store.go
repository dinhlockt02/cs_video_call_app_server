package authmiddleware

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *MongoStore {
	return &MongoStore{database: database}
}

func (s *MongoStore) FindOne(ctx context.Context, filter map[string]interface{}) (*Device, error) {
	var device Device
	result := s.database.Collection(device.CollectionName()).FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "can not find device")
	}
	if err := result.Decode(&device); err != nil {
		return nil, errors.Wrap(err, "can not decode device")
	}
	return &device, nil
}
