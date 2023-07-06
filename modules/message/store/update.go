package messagestore

import (
	"context"
	messagemdl "github.com/dinhlockt02/cs_video_call_app_server/modules/message/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *MongoStore) UpdateMany(ctx context.Context, filter map[string]interface{},
	data *messagemdl.UpdateMessage) error {
	log.Debug().Any("filter", filter).Any("data", data).Msg("update messages")
	update := bson.M{"$set": data}

	_, err := s.database.Collection(data.CollectionName()).UpdateMany(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "can not update messages")
	}
	return nil
}
