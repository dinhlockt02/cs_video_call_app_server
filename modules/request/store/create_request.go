package requeststore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) CreateRequest(ctx context.Context, request *requestmdl.Request) error {
	result, err := s.database.Collection(request.CollectionName()).InsertOne(ctx, request)
	if err != nil {
		return common.ErrInternal(err)
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	request.Id = &insertedId
	return nil
}
