package meetingrepo

import (
	"context"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	meetingstore "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/store"
	userstore "github.com/dinhlockt02/cs_video_call_app_server/modules/user/store"
)

type Repository interface {
	CreateMeeting(ctx context.Context, meeting *meetingmodel.Meeting) error
	ListMeeting(ctx context.Context, filter map[string]interface{}) ([]meetingmodel.Meeting, error)
	FindMeeting(ctx context.Context, filter map[string]interface{}) (*meetingmodel.Meeting, error)
	FindParticipant(ctx context.Context, filter map[string]interface{}) (*meetingmodel.Participant, error)
	UpdateMeeting(ctx context.Context, filter map[string]interface{}, data *meetingmodel.UpdateMeeting) error
}

type meetingRepository struct {
	meetingStore meetingstore.Store
	userStore    userstore.Store
}

func NewMeetingRepository(
	meetingStore meetingstore.Store,
	userStore userstore.Store,
) Repository {
	return &meetingRepository{
		meetingStore: meetingStore,
		userStore:    userStore,
	}
}
