package meetingbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	lksv "github.com/dinhlockt02/cs_video_call_app_server/components/livekit_service"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	meetingrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/repository"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type CreateMeetingBiz struct {
	meetingRepo    meetingrepo.Repository
	livekitService lksv.LiveKitService
}

func NewCreateMeetingBiz(
	meetingRepo meetingrepo.Repository,
	livekitService lksv.LiveKitService,
) *CreateMeetingBiz {
	return &CreateMeetingBiz{
		meetingRepo:    meetingRepo,
		livekitService: livekitService,
	}
}

func (biz *CreateMeetingBiz) Create(ctx context.Context,
	requesterId string, meeting *meetingmodel.Meeting) (string, error) {
	log.Debug().Str("requesterId", requesterId).Any("meeting", meeting).Msg("create meeting")

	// Find requester
	requesterFilter, err := common.GetIdFilter(requesterId)
	if err != nil {
		return "", common.ErrInvalidRequest(errors.Wrap(err, "invalid requester id"))
	}

	requester, err := biz.meetingRepo.FindParticipant(ctx, requesterFilter)

	if err != nil {
		return "", common.ErrInternal(errors.Wrap(err, "can not find requester"))
	}

	if requester == nil {
		return "", common.ErrEntityNotFound(common.UserEntity, errors.New("requester not found"))
	}

	// Create meeting
	meeting.Status = meetingmodel.OnGoing
	meeting.Participants = append(meeting.Participants, *requester)

	err = biz.meetingRepo.CreateMeeting(ctx, meeting)
	if err != nil {
		return "", common.ErrInternal(errors.Wrap(err, "can not create meeting"))
	}
	_, err = biz.livekitService.CreateRoom(ctx, *meeting.Id)
	if err != nil {
		return "", common.ErrInternal(errors.Wrap(err, "can not create livekit room"))
	}

	token, err := biz.livekitService.CreateJoinToken(*meeting.Id, requesterId)
	if err != nil {
		return "", common.ErrInternal(errors.Wrap(err, "can not create join token"))
	}
	return token, nil
}
