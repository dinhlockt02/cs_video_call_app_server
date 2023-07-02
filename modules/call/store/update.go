package callstore

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) Update(ctx context.Context, filter map[string]interface{}, data *callmdl.Call) error {
	id := data.Id
	data.Id = nil
	update := bson.M{"$set": data}

	_, err := s.database.Collection(data.CollectionName()).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	data.Id = id
	return nil
}
