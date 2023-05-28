package callstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) FindCall(ctx context.Context, filter map[string]interface{}) (*callmdl.Call, error) {
	var call callmdl.Call
	result := s.database.Collection(call.CollectionName()).FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}

	if err := result.Decode(&call); err != nil {
		return nil, common.ErrInternal(err)
	}
	return &call, nil
}
