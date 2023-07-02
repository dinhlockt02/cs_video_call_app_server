package meetingstore

import (
	"context"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *mongoStore) ListMeeting(ctx context.Context, filter map[string]interface{}) ([]meetingmodel.Meeting, error) {
	log.Debug().Any("filter", filter).Msg("list meetings")

	var meetings []meetingmodel.Meeting

	findOptions := options.Find()
	// Sort by `price` field descending
	findOptions.SetSort(bson.M{"time_start": -1})

	cursor, err := s.database.
		Collection(meetingmodel.Meeting{}.CollectionName()).
		Find(ctx, filter, findOptions)
	if err != nil {
		return nil, errors.Wrap(err, "can not list meetings")
	}
	err = cursor.All(ctx, &meetings)
	if err != nil {
		return nil, errors.Wrap(err, "can not decode meetings")
	}
	return meetings, nil
}
