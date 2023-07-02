package notistore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MongoStore) Create(ctx context.Context, data *notimodel.Notification) error {
	result, err := s.database.Collection(data.CollectionName()).InsertOne(ctx, data)
	if err != nil {
		return common.ErrInternal(err)
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	data.Id = &insertedId
	return nil
}
