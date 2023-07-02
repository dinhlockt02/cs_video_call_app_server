package authstore

import (
	"context"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MongoStore) Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error) {
	log.Debug().Any("filter", filter).Msg("find a user")
	var findUser authmodel.User
	result := s.database.Collection(findUser.CollectionName()).FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "can not find user")
	}
	err := result.Decode(&findUser)
	if err != nil {
		return nil, errors.Wrap(err, "can not decode user")
	}
	log.Debug().Any("filter", filter).Msg("a user found")
	return &findUser, nil
}
