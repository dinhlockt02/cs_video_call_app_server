package meetingstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *mongoStore) ListMeeting(ctx context.Context, filter map[string]interface{}) ([]meetingmodel.Meeting, error) {
	var meetings []meetingmodel.Meeting

	findOptions := options.Find()
	// Sort by `price` field descending
	findOptions.SetSort(bson.D{{"time_start", -1}})

	cursor, err := s.database.
		Collection(meetingmodel.Meeting{}.CollectionName()).
		Find(ctx, filter, findOptions)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	err = cursor.All(ctx, &meetings)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return meetings, nil
}
