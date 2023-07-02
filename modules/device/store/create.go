package devicestore

import (
	"context"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) Create(ctx context.Context, data *devicemodel.Device) error {
	log.Debug().Any("data", data).Msg("create a device")
	result, err := s.database.Collection(data.CollectionName()).InsertOne(ctx, data)
	if err != nil {
		return errors.Wrap(err, "can not insert device")
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	data.Id = &id
	return nil
}
