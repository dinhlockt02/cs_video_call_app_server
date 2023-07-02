package groupstore

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *mongoStore) DeleteOne(ctx context.Context, filter map[string]interface{}) error {
	log.Debug().Any("filter", filter).Msg("delete group")
	_, err := s.database.Collection(groupmdl.Group{}.CollectionName()).DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "can not delete group")
	}
	return nil
}
