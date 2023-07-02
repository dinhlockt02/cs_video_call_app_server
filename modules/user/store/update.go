package userstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *MongoStore) Update(ctx context.Context,
	filter map[string]interface{}, updatedUser *usermodel.UpdateUser) error {
	log.Debug().Any("filter", filter).Any("updatedUser", updatedUser).Msg("update user")
	update := bson.M{"$set": updatedUser}

	_, err := s.database.
		Collection(updatedUser.CollectionName()).
		UpdateOne(ctx, filter, update)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
