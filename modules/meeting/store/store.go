package meetingstore

import (
	"context"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	CreateMeeting(ctx context.Context, meeting *meetingmodel.Meeting) error
	UpdateMeeting(
		ctx context.Context,
		filter map[string]interface{},
		updatedMeeting *meetingmodel.UpdateMeeting,
	) error
	ListMeeting(ctx context.Context, filter map[string]interface{}) ([]meetingmodel.Meeting, error)
	FindMeeting(ctx context.Context, filter map[string]interface{}) (*meetingmodel.Meeting, error)
}

type mongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *mongoStore {
	return &mongoStore{database: database}
}
