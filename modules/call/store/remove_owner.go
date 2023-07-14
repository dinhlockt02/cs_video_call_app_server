package callstore

import (
	"context"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *MongoStore) RemoveOwner(ctx context.Context, filter map[string]interface{}, owner string) error {
	log.Debug().Any("filter", filter).Any("owner", owner).Msg("Remove owner")
	update := bson.M{"$pull": bson.M{"owner": owner}}

	_, err := s.database.Collection(callmdl.Call{}.CollectionName()).UpdateMany(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "can not update call")
	}
	return nil
}
