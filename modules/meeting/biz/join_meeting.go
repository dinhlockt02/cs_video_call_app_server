package meetingbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	lksv "github.com/dinhlockt02/cs_video_call_app_server/components/livekit_service"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	meetingrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/repository"
	meetingstore "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/store"
)

type joinMeetingBiz struct {
	meetingRepo    meetingrepo.Repository
	livekitService lksv.LiveKitService
}

func NewJoinMeetingBiz(
	meetingRepo meetingrepo.Repository,
	livekitService lksv.LiveKitService,
) *joinMeetingBiz {
	return &joinMeetingBiz{
		meetingRepo:    meetingRepo,
		livekitService: livekitService,
	}
}

func (biz *joinMeetingBiz) Join(ctx context.Context, requester, groupId, meetingId string) (string, error) {
	// Create meeting

	idFilter := map[string]interface{}{}
	err := common.AddIdFilter(idFilter, meetingId)
	if err != nil {
		return "", err
	}

	meeting, err := biz.meetingRepo.FindMeeting(ctx, common.GetAndFilter(idFilter, meetingstore.GetGroupFilter(groupId)))
	if err != nil {
		return "", err
	}
	if meeting == nil {
		return "", common.ErrEntityNotFound("Meeting", meetingmodel.ErrMeetingNotFound)
	}

	if meeting.Status == meetingmodel.Ended {
		return "", common.ErrInvalidRequest(meetingmodel.ErrMeetingEnded)

	}

	token, err := biz.livekitService.CreateJoinToken(*meeting.Id, requester)
	if err != nil {
		return "", common.ErrInternal(err)

	}
	return token, nil
}
