package devicestore

import (
	"context"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *mongoStore) Delete(ctx context.Context, filter map[string]interface{}) error {
	log.Debug().Any("filter", filter).Msg("delete many devices")
	result, err := s.database.Collection(devicemodel.Device{}.CollectionName()).DeleteMany(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "can not delete devices")
	}
	log.Debug().Int64("DeleteCount", result.DeletedCount).Msg("deleted")
	return nil
}
