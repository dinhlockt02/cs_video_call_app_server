package notistore

import (
	"context"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *MongoStore) Delete(ctx context.Context, filter map[string]interface{}) error {
	log.Debug().Any("filter", filter).Msg("delete notifications")
	_, err := s.database.Collection(notimodel.Notification{}.CollectionName()).DeleteMany(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "can not delete notifications")
	}
	return nil
}
