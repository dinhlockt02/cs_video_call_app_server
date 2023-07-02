package requeststore

import (
	"context"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *mongoStore) DeleteRequest(ctx context.Context, filter map[string]interface{}) error {
	log.Debug().Any("filter", filter).Msg("delete request")
	_, err := s.database.Collection((&requestmdl.Request{}).CollectionName()).DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "can not delete request")
	}
	return nil
}
