package callstore

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) Update(ctx context.Context, filter map[string]interface{}, data *callmdl.Call) error {
	log.Debug().Any("filter", filter).Any("data", data).Msg("update a call")
	id := data.Id
	data.Id = nil
	update := bson.M{"$set": data}

	_, err := s.database.Collection(data.CollectionName()).UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "can not update call")
	}
	data.Id = id
	return nil
}
