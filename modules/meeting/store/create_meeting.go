package meetingstore

import (
	"context"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MongoStore) CreateMeeting(ctx context.Context, meeting *meetingmodel.Meeting) error {
	log.Debug().Any("meeting", meeting).Msg("create meeting")
	result, err := s.database.Collection(meeting.CollectionName()).InsertOne(ctx, meeting)
	if err != nil {
		return errors.Wrap(err, "can not create meeting")
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	meeting.Id = &insertedId
	return nil
}
