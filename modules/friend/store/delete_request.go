package friendstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) DeleteRequest(ctx context.Context, requestId string) error {
	id, _ := primitive.ObjectIDFromHex(requestId)
	_, err := s.database.Collection(friendmodel.Request{}.CollectionName()).DeleteOne(ctx, bson.D{
		{"_id", id},
	})
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
