package groupstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) FindUser(ctx context.Context, filter map[string]interface{}) (*groupmdl.User, error) {
	var user groupmdl.User
	result := s.database.Collection(user.CollectionName()).FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}

	if err := result.Decode(&user); err != nil {
		return nil, common.ErrInternal(err)
	}

	return &user, nil
}
