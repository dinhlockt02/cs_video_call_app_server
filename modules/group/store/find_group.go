package groupstore

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MongoStore) FindGroup(
	ctx context.Context,
	filter map[string]interface{},
) (*groupmdl.Group, error) {
	log.Debug().Any("filter", filter).Msg("find a group")
	var group *groupmdl.Group

	result := s.database.Collection(groupmdl.Group{}.CollectionName()).
		FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "can not find group")
	}
	if err := result.Decode(&group); err != nil {
		return nil, errors.Wrap(err, "can not find group")
	}
	return group, nil
}
