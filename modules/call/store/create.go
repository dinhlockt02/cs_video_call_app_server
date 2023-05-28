package callstore

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) Create(ctx context.Context, data *callmdl.Call) error {
	result, err := s.database.Collection(data.CollectionName()).InsertOne(ctx, data)
	if err != nil {
		return err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	data.Id = &id
	return nil
}
