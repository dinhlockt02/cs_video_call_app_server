package groupstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
)

func (s *mongoStore) FindUsers(ctx context.Context, filter map[string]interface{}) ([]groupmdl.User, error) {
	var users []groupmdl.User
	cursor, err := s.database.Collection(groupmdl.User{}.CollectionName()).Find(ctx, filter)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if err = cursor.All(ctx, &users); err != nil {
		return nil, common.ErrInternal(err)
	}

	return users, nil
}
