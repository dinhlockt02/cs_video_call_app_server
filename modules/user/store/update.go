package userstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) Update(ctx context.Context, filter map[string]interface{}, updatedUser *usermodel.UpdateUser) error {
	update := bson.E{Key: "$set", Value: updatedUser}

	_, err := s.database.
		Collection(updatedUser.CollectionName()).
		UpdateOne(ctx, filter, update)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
