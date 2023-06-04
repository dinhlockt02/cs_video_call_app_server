package meetingrepo

import (
	"context"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
)

func (r *meetingRepository) ListMeeting(ctx context.Context, filter map[string]interface{}) ([]meetingmodel.Meeting, error) {
	return r.meetingStore.ListMeeting(ctx, filter)
}
