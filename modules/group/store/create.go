package groupstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) Create(ctx context.Context, group *groupmdl.Group) error {
	result, err := s.database.Collection(group.CollectionName()).InsertOne(ctx, group)
	if err != nil {
		return common.ErrInternal(err)
	}
	createdId := result.InsertedID.(primitive.ObjectID).Hex()
	group.Id = &createdId
	return nil
}
