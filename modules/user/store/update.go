package userstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) Update(ctx context.Context, updatedUser *usermodel.UpdateUser) error {

	update := bson.D{{"$set", updatedUser}}

	_, err := s.database.
		Collection(updatedUser.CollectionName()).
		UpdateByID(ctx, updatedUser.Id, update)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
