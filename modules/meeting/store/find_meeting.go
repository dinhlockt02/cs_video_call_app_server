package meetingstore

import (
	"context"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MongoStore) FindMeeting(ctx context.Context, filter map[string]interface{}) (*meetingmodel.Meeting, error) {
	log.Debug().Any("filter", filter).Msg("find meeting")

	var meeting meetingmodel.Meeting

	result := s.database.
		Collection(meetingmodel.Meeting{}.CollectionName()).
		FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "can not find meeting")
	}

	if err := result.Decode(&meeting); err != nil {
		return nil, errors.Wrap(err, "can not decode meeting")
	}
	return &meeting, nil
}
