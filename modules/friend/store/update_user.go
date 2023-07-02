package friendstore

import (
	"context"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MongoStore) UpdateUser(ctx context.Context,
	filter map[string]interface{}, updatedUser *friendmodel.User) error {
	log.Debug().Any("filter", filter).Any("updatedUser", updatedUser).Msg("update a user")
	updatedUser.Id = nil
	updateData := bson.M{
		"$set": updatedUser,
	}
	_, err := s.database.
		Collection(updatedUser.CollectionName()).
		UpdateOne(ctx, filter, updateData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return errors.Wrap(err, "can not update user")
	}
	return nil
}
