package authstore

import (
	"context"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) Update(ctx context.Context, filter map[string]interface{}, passwordUser *authmodel.UpdatePasswordUser) error {
	log.Debug().Any("filter", filter).Any("passwordUser", passwordUser).Msg("update user")
	update := bson.M{"$set": passwordUser}

	_, err := s.database.
		Collection(passwordUser.CollectionName()).
		UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "can not update user")
	}
	log.Debug().Any("filter", filter).Any("passwordUser", passwordUser).Msg("user updated")
	return nil
}
