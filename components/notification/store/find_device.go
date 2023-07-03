package notistore

import (
	"context"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *MongoStore) FindDevice(ctx context.Context, filter map[string]interface{}) ([]notimodel.Device, error) {
	log.Debug().Any("filter", filter).Msg("find devices")
	cursor, err := s.database.Collection(notimodel.Device{}.CollectionName()).
		Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find devices")
	}
	var devices []notimodel.Device
	if err = cursor.All(ctx, &devices); err != nil {
		return nil, errors.Wrap(err, "can not decode devices")
	}
	return devices, nil
}
