package meetingstore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) UpdateMeeting(
	ctx context.Context,
	filter map[string]interface{},
	updatedMeeting *meetingmodel.UpdateMeeting,
) error {
	data := bson.D{{"$set", updatedMeeting}}
	_, err := s.database.Collection(updatedMeeting.CollectionName()).UpdateOne(ctx, filter, data)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
