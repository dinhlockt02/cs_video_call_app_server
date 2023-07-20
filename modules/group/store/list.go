package groupstore

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *MongoStore) List(
	ctx context.Context,
	filter map[string]interface{},
) ([]groupmdl.Group, error) {
	log.Debug().Any("filter", filter).Msg("list groups")

	cursor, err := s.database.Collection(groupmdl.Group{}.CollectionName()).
		Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "can not list groups")
	}
	var groups []groupmdl.Group
	if err = cursor.All(ctx, &groups); err != nil {
		return nil, errors.Wrap(err, "can not decode groups")
	}
	return groups, nil
}
