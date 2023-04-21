package userstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) Update(ctx context.Context, updatedUser *usermodel.UpdateUser) error {

	update := bson.D{{"$set", updatedUser}}

	primitiveId, err := primitive.ObjectIDFromHex(*updatedUser.Id)
	updatedUser.Id = nil
	if err != nil {
		return common.ErrInvalidRequest(err)
	}

	_, err = s.database.
		Collection(updatedUser.CollectionName()).
		UpdateByID(ctx, primitiveId, update)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
