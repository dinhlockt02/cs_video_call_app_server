package userstore

import (
	"context"
	usermodel "github.com/dinhlockt02/cs_video_call_app_server/modules/user/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error) {
	log.Debug().Any("filter", filter).Msg("find user")
	result := s.database.
		Collection(usermodel.User{}.CollectionName()).
		FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "can not find one user")
	}
	var user = new(usermodel.User)
	if err := result.Decode(&user); err != nil {
		return nil, errors.Wrap(err, "can not decode user")
	}
	return user, nil
}
