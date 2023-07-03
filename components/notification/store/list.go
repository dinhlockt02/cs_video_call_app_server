package notistore

import (
	"context"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *MongoStore) List(ctx context.Context, filter map[string]interface{}) ([]notimodel.Notification, error) {
	log.Debug().Any("filter", filter).Msg("list notifications")
	opts := options.Find().SetSort(bson.M{"created_at": -1})
	cursor, err := s.database.Collection(notimodel.Notification{}.CollectionName()).Find(ctx, filter, opts)
	if err != nil {
		return nil, errors.Wrap(err, "can not find notifications")
	}
	var result []notimodel.Notification
	if err = cursor.All(ctx, &result); err != nil {
		log.Error().Err(err).Str("pakage", "notistore.Find").Send()
		return nil, errors.Wrap(err, "can not decode notifications")
	}

	return result, nil
}
