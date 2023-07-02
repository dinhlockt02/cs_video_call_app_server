package groupstore

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MongoStore) Create(ctx context.Context, group *groupmdl.Group) error {
	log.Debug().Any("group", group).Msg("create group")
	result, err := s.database.Collection(group.CollectionName()).InsertOne(ctx, group)
	if err != nil {
		return errors.Wrap(err, "can not create group")
	}
	createdId := result.InsertedID.(primitive.ObjectID).Hex()
	group.Id = &createdId
	return nil
}
