package meetingstore

import (
	"context"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *MongoStore) UpdateParticipants(
	ctx context.Context,
	filter map[string]interface{},
	updateParticipant *meetingmodel.Participant,
) error {
	log.Debug().Any("filter", filter).Any("updateParticipant", updateParticipant).Msg("list meetings")
	data := bson.M{"$set": bson.M{"participants.$": updateParticipant}}
	_, err := s.database.Collection((&meetingmodel.Meeting{}).CollectionName()).UpdateMany(ctx, filter, data)
	if err != nil {
		return errors.Wrap(err, "can not update meeting")
	}
	return nil
}
