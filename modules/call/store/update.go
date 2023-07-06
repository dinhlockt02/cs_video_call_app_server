package callstore

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *MongoStore) Update(ctx context.Context, filter map[string]interface{}, data *callmdl.UpdateCall) error {
	log.Debug().Any("filter", filter).Any("data", data).Msg("update a call")
	update := bson.M{"$set": data}

	_, err := s.database.Collection(data.CollectionName()).UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "can not update call")
	}
	return nil
}

func (s *MongoStore) UpdateMany(ctx context.Context, filter map[string]interface{}, data *callmdl.UpdateCall) error {
	log.Debug().Any("filter", filter).Any("data", data).Msg("update a call")
	update := bson.M{"$set": data}

	_, err := s.database.Collection(data.CollectionName()).UpdateMany(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "can not update call")
	}
	return nil
}
