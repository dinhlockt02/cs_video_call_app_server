package authstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
)

func (s *mongoStore) DeleteUser(ctx context.Context, filter map[string]interface{}) error {
	_, err := s.database.Collection(authmodel.User{}.CollectionName()).DeleteOne(ctx, filter)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
