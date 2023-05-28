package groupstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
)

func (s *mongoStore) List(
	ctx context.Context,
	filter map[string]interface{},
) ([]groupmdl.Group, error) {

	cursor, err := s.database.Collection(groupmdl.Group{}.CollectionName()).
		Find(ctx, filter)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	var groups []groupmdl.Group
	if err = cursor.All(ctx, &groups); err != nil {
		return nil, common.ErrInternal(err)
	}
	return groups, nil
}
