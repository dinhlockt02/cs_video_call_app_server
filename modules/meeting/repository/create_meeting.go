package meetingrepo

import (
	"context"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
)

func (r *meetingRepository) CreateMeeting(ctx context.Context, meeting *meetingmodel.Meeting) error {
	return r.meetingStore.CreateMeeting(ctx, meeting)
}
