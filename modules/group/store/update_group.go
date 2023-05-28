package groupstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) UpdateGroup(
	ctx context.Context,
	filter map[string]interface{},
	updatedGroup *groupmdl.Group,
) error {
	updatedGroup.Id = nil
	updateData := bson.D{{
		"$set", updatedGroup,
	}}
	_, err := s.database.
		Collection(updatedGroup.CollectionName()).
		UpdateOne(ctx, filter, updateData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return common.ErrInternal(err)
	}
	return nil
}
