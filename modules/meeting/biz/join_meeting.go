package meetingbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	lksv "github.com/dinhlockt02/cs_video_call_app_server/components/livekit_service"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	meetingrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/repository"
	meetingstore "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type JoinMeetingBiz struct {
	meetingRepo    meetingrepo.Repository
	livekitService lksv.LiveKitService
}

func NewJoinMeetingBiz(
	meetingRepo meetingrepo.Repository,
	livekitService lksv.LiveKitService,
) *JoinMeetingBiz {
	return &JoinMeetingBiz{
		meetingRepo:    meetingRepo,
		livekitService: livekitService,
	}
}

func (biz *JoinMeetingBiz) Join(ctx context.Context, requester, groupId, meetingId string) (string, error) {
	log.Debug().Str("requester", requester).
		Str("groupId", groupId).
		Str("meetingId", meetingId).
		Msg("join meeting")

	idFilter, err := common.GetIdFilter(meetingId)
	if err != nil {
		return "", common.ErrInvalidRequest(errors.Wrap(err, "invalid meeting id"))
	}

	meeting, err := biz.meetingRepo.FindMeeting(ctx, common.GetAndFilter(idFilter, meetingstore.GetGroupFilter(groupId)))
	if err != nil {
		return "", common.ErrInternal(errors.Wrap(err, "can not find meeting"))
	}
	if meeting == nil {
		return "", common.ErrEntityNotFound(common.MeetingEntity, errors.New(meetingmodel.MeetingNotFound))
	}

	if meeting.Status == meetingmodel.Ended {
		return "", common.ErrInvalidRequest(errors.New(meetingmodel.MeetingEnded))
	}

	token, err := biz.livekitService.CreateJoinToken(*meeting.Id, requester)
	if err != nil {
		return "", common.ErrInternal(errors.Wrap(err, "can not create join room token"))
	}
	return token, nil
}
