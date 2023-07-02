package friendstore

import (
	"context"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MongoStore) FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error) {
	log.Debug().Any("filter", filter).Msg("find a user")
	var user friendmodel.User
	result := s.database.Collection(user.CollectionName()).FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "can not find a user")
	}

	if err := result.Decode(&user); err != nil {
		return nil, errors.Wrap(err, "can not decode user data")
	}

	return &user, nil
}
