package notistore

import (
	"context"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *MongoStore) UpdateUser(ctx context.Context,
	filter map[string]interface{},
	updatedData *notimodel.NotificationUser) error {
	log.Debug().
		Any("filter", filter).
		Any("updatedData", updatedData).
		Msg("update user's notification setting")
	_, err := s.database.
		Collection(notimodel.NotificationUser{}.CollectionName()).
		UpdateOne(ctx, filter, bson.M{"$set": updatedData})
	if err != nil {
		return errors.Wrap(err, "can not update user notification setting")
	}
	return nil
}
