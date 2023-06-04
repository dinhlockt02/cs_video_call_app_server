package meetingstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) FindMeeting(ctx context.Context, filter map[string]interface{}) (*meetingmodel.Meeting, error) {
	var meeting meetingmodel.Meeting

	result := s.database.
		Collection(meetingmodel.Meeting{}.CollectionName()).
		FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}

	if err := result.Decode(&meeting); err != nil {
		return nil, common.ErrInternal(err)
	}
	return &meeting, nil
}
