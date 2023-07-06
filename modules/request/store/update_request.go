package requeststore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *MongoStore) UpdateRequests(
	ctx context.Context,
	filter map[string]interface{},
	data *requestmdl.UpdateRequest,
) error {
	updateData := bson.D{{
		"$set", data,
	}}
	_, err := s.database.Collection(data.CollectionName()).UpdateMany(ctx, filter, updateData)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
