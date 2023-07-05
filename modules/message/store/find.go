package messagestore

import (
	"context"
	messagemdl "github.com/dinhlockt02/cs_video_call_app_server/modules/message/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *MongoStore) List(ctx context.Context, filter map[string]interface{}) ([]messagemdl.Message, error) {
	log.Debug().Any("filter", filter).Msg("list messages")
	var messages []messagemdl.Message

	opt := options.Find().SetSort(bson.M{"created_at": 1})
	cursor, err := s.database.Collection((&messagemdl.Message{}).CollectionName()).Find(ctx, filter, opt)
	if err != nil {
		return nil, errors.Wrap(err, "can not list messages")
	}

	if err = cursor.All(ctx, &messages); err != nil {
		return nil, errors.Wrap(err, "can not decode messages")
	}

	return messages, nil
}

func (s *MongoStore) Get(ctx context.Context, filter map[string]interface{}) (*messagemdl.Message, error) {
	log.Debug().Any("filter", filter).Msg("find a requests")
	var message messagemdl.Message
	result := s.database.Collection(message.CollectionName()).FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "can not find a message")
	}

	if err := result.Decode(&message); err != nil {
		return nil, errors.Wrap(err, "can not decode message")
	}
	return &message, nil
}
