package messagestore

import (
	"context"
	messagemdl "github.com/dinhlockt02/cs_video_call_app_server/modules/message/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MongoStore) Create(ctx context.Context, message *messagemdl.Message) error {
	log.Debug().Any("message", message).Msg("create message")
	result, err := s.database.Collection(message.CollectionName()).InsertOne(ctx, message)
	if err != nil {
		return errors.Wrap(err, "can not create request")
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	message.Id = &insertedId
	return nil
}
