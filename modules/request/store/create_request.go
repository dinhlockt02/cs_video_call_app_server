package requeststore

import (
	"context"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) CreateRequest(ctx context.Context, request *requestmdl.Request) error {
	log.Debug().Any("request", request).Msg("create request")
	result, err := s.database.Collection(request.CollectionName()).InsertOne(ctx, request)
	if err != nil {
		return errors.Wrap(err, "can not create request")
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	request.Id = &insertedId
	return nil
}
