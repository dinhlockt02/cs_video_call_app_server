package messagestore

import (
	"context"
	messagemdl "github.com/dinhlockt02/cs_video_call_app_server/modules/message/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *MongoStore) DeleteOne(ctx context.Context, filter map[string]interface{}) error {
	log.Debug().Any("filter", filter).Msg("delete message")
	_, err := s.database.Collection((&messagemdl.Message{}).CollectionName()).DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "can not delete message")
	}
	return nil
}
