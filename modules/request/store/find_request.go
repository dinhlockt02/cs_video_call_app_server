package requeststore

import (
	"context"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MongoStore) FindRequests(ctx context.Context, filter map[string]interface{}) ([]requestmdl.Request, error) {
	log.Debug().Any("filter", filter).Msg("find requests")
	var request []requestmdl.Request
	cursor, err := s.database.Collection((&requestmdl.Request{}).CollectionName()).Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not find requests")
	}

	if err = cursor.All(ctx, &request); err != nil {
		return nil, errors.Wrap(err, "can not decode requests")
	}

	return request, nil
}

func (s *MongoStore) FindRequest(ctx context.Context, filter map[string]interface{}) (*requestmdl.Request, error) {
	log.Debug().Any("filter", filter).Msg("find a requests")
	var request requestmdl.Request
	result := s.database.Collection(request.CollectionName()).FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "can not find a request")
	}

	if err := result.Decode(&request); err != nil {
		return nil, errors.Wrap(err, "can not decode request")
	}
	return &request, nil
}
