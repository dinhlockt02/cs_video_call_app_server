package authstore

import (
	"context"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *MongoStore) UpdateEmailVerified(ctx context.Context, filter map[string]interface{}) error {
	log.Debug().Any("filter", filter).Msg("update user's email verify status")
	// TODO: move this up to biz
	var updateEmailVerifiedUser authmodel.EmailVerifiedUser
	updateEmailVerifiedUser.Process()

	update := bson.M{"$set": updateEmailVerifiedUser}

	_, err := s.database.
		Collection(updateEmailVerifiedUser.CollectionName()).
		UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "can not update user")
	}
	log.Debug().Any("filter", filter).Msg("updated user's email verify status")
	return nil
}
