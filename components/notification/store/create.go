package notistore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MongoStore) Create(ctx context.Context, data *notimodel.Notification) error {
	log.Debug().Any("data", data).Msg("create a notification")
	result, err := s.database.Collection(data.CollectionName()).InsertOne(ctx, data)
	if err != nil {
		return errors.Wrap(err, "can not create notification")
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	data.Id = &insertedId
	return nil
}
