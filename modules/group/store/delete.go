package groupstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
)

func (s *mongoStore) DeleteOne(ctx context.Context, filter map[string]interface{}) error {
	_, err := s.database.Collection(groupmdl.Group{}.CollectionName()).DeleteOne(ctx, filter)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
