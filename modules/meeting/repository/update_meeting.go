package meetingrepo

import (
	"context"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
)

func (r *meetingRepository) UpdateMeeting(ctx context.Context,
	filter map[string]interface{}, data *meetingmodel.UpdateMeeting) error {
	return r.meetingStore.UpdateMeeting(ctx, filter, data)
}
