package devicestore

import (
	"context"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *MongoStore) Get(ctx context.Context, filter map[string]interface{}) ([]*devicemodel.GetDeviceDto, error) {
	log.Debug().Any("filter", filter).Msg("find devices")
	var devices []*devicemodel.GetDeviceDto
	cursor, err := s.database.Collection(devicemodel.GetDeviceDto{}.CollectionName()).Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not get devices")
	}
	if err = cursor.All(ctx, &devices); err != nil {
		return nil, errors.Wrap(err, "can not decode device documents to GetDeviceDto")
	}
	return devices, nil
}
