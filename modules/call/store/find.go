package callstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MongoStore) FindCall(ctx context.Context, filter map[string]interface{}) (*callmdl.Call, error) {
	log.Debug().Any("filter", filter).Msg("find a call")
	var call callmdl.Call
	result := s.database.Collection(call.CollectionName()).FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "can not find call")
	}

	if err := result.Decode(&call); err != nil {
		return nil, common.ErrInternal(err)
	}
	return &call, nil
}
