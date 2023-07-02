package groupstore

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) UpdateGroup(
	ctx context.Context,
	filter map[string]interface{},
	updatedGroup *groupmdl.Group,
) error {
	log.Debug().Any("filter", filter).Any("updatedGroup", updatedGroup).Msg("update a group")
	updatedGroup.Id = nil
	updateData := bson.M{
		"$set": updatedGroup,
	}
	_, err := s.database.
		Collection(updatedGroup.CollectionName()).
		UpdateOne(ctx, filter, updateData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return errors.Wrap(err, "can not update group")
	}
	return nil
}
