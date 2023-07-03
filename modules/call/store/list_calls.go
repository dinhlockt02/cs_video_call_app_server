package callstore

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *MongoStore) ListCalls(ctx context.Context, filter map[string]interface{}) ([]callmdl.Call, error) {
	log.Debug().Any("filter", filter).Msg("list calls")
	opt := options.Find().SetSort(bson.M{"called_at": -1})
	result, err := s.database.Collection((&callmdl.Call{}).CollectionName()).Find(ctx, filter, opt)
	if err != nil {
		return nil, errors.Wrap(err, "can not list calls")
	}
	var calls []callmdl.Call
	if err = result.All(ctx, &calls); err != nil {
		return nil, errors.Wrap(err, "can not decode calls")
	}

	return calls, nil
}
