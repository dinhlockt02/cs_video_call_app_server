package requeststore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) FindRequests(ctx context.Context, filter map[string]interface{}) ([]requestmdl.Request, error) {
	var request []requestmdl.Request
	log.Info().Msgf("%v", filter)
	cursor, err := s.database.Collection(requestmdl.Request{}.CollectionName()).Find(ctx, filter)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if err = cursor.All(ctx, &request); err != nil {
		return nil, common.ErrInternal(err)
	}

	return request, nil
}

func (s *mongoStore) FindRequest(ctx context.Context, filter map[string]interface{}) (*requestmdl.Request, error) {
	var request requestmdl.Request
	result := s.database.Collection(request.CollectionName()).FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}

	if err := result.Decode(&request); err != nil {
		return nil, common.ErrInternal(err)
	}
	return &request, nil
}
