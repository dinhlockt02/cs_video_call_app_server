package meetingstore

import (
	"context"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *MongoStore) UpdateMeeting(
	ctx context.Context,
	filter map[string]interface{},
	updatedMeeting *meetingmodel.UpdateMeeting,
) error {
	log.Debug().Any("filter", filter).Any("updatedMeeting", updatedMeeting).Msg("list meetings")
	data := bson.M{"$set": updatedMeeting}
	_, err := s.database.Collection(updatedMeeting.CollectionName()).UpdateOne(ctx, filter, data)
	if err != nil {
		return errors.Wrap(err, "can not update meeting")
	}
	return nil
}
