package notistore

import (
	"context"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MongoStore) FindUser(ctx context.Context,
	filter map[string]interface{}) (*notimodel.NotificationUser, error) {
	result := s.database.
		Collection(notimodel.NotificationUser{}.CollectionName()).
		FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "can not find notification user")
	}
	var user notimodel.NotificationUser
	if err := result.Decode(&user); err != nil {
		return nil, errors.Wrap(err, "can not decode user")
	}
	return &user, nil
}
