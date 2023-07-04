package authstore

import (
	"context"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *MongoStore) ResetPassword(ctx context.Context, filter map[string]interface{}, data *authmodel.ResetPasswordBody) error {
	update := bson.D{{"$set", data}}

	_, err := s.database.
		Collection(data.CollectionName()).
		UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "can not reset password")
	}
	return nil
}
