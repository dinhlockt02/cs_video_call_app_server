package authstore

import (
	"context"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *mongoStore) DeleteUser(ctx context.Context, filter map[string]interface{}) error {
	log.Debug().Any("filter", filter).Msg("deleting a user")
	rs, err := s.database.Collection(authmodel.User{}.CollectionName()).DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "can not delete user")
	}
	log.Debug().Any("filter", filter).Int64("count", rs.DeletedCount).Msg("delete a user successful")
	return nil
}
