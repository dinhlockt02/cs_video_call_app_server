package notistore

import (
	"context"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *MongoStore) UpdateNotifications(
	ctx context.Context,
	filter map[string]interface{},
	data *notimodel.UpdateNotification,
) error {
	log.Debug().Any("filter", filter).Any("data", data).Msg("update notifications")
	updateData := bson.M{
		"$set": data,
	}
	_, err := s.database.Collection(data.CollectionName()).UpdateMany(ctx, filter, updateData)
	if err != nil {
		return errors.Wrap(err, "can not update notifications")
	}
	return nil
}
