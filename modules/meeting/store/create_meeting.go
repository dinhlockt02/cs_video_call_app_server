package meetingstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) CreateMeeting(ctx context.Context, meeting *meetingmodel.Meeting) error {
	result, err := s.database.Collection(meeting.CollectionName()).InsertOne(ctx, meeting)
	if err != nil {
		return common.ErrInternal(err)
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	meeting.Id = &insertedId
	return nil
}
