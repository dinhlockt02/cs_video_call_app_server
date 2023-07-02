package meetingbiz

import (
	"context"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	meetingrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/repository"
	meetingstore "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/store"
	"github.com/rs/zerolog/log"
)

type listMeetingsBiz struct {
	meetingRepo meetingrepo.Repository
}

func NewListMeetingsBiz(
	meetingRepo meetingrepo.Repository,
) *listMeetingsBiz {
	return &listMeetingsBiz{
		meetingRepo: meetingRepo,
	}
}

func (biz *listMeetingsBiz) List(ctx context.Context, groupId string) ([]meetingmodel.Meeting, error) {
	log.Debug().Str("groupId", groupId).
		Msg("list meetings")
	filter := meetingstore.GetGroupFilter(groupId)
	return biz.meetingRepo.ListMeeting(ctx, filter)
}
