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

type createMeetingBiz struct {
	meetingRepo    meetingrepo.Repository
	livekitService lksv.LiveKitService
}

func NewCreateMeetingBiz(
	meetingRepo meetingrepo.Repository,
	livekitService lksv.LiveKitService,
) *createMeetingBiz {
	return &createMeetingBiz{
		meetingRepo:    meetingRepo,
		livekitService: livekitService,
	}
}

func (biz *createMeetingBiz) Create(ctx context.Context, requester string, meeting *meetingmodel.Meeting) (string, error) {
	log.Debug().Str("requester", requester).Any("meeting", meeting).Msg("create meeting")
	// Create meeting
	meeting.Status = meetingmodel.OnGoing

	err := biz.meetingRepo.CreateMeeting(ctx, meeting)
	if err != nil {
		return "", common.ErrInternal(errors.Wrap(err, "can not create meeting"))
	}
	_, err = biz.livekitService.CreateRoom(ctx, *meeting.Id)
	if err != nil {
		return "", common.ErrInternal(errors.Wrap(err, "can not create livekit room"))
	}

	token, err := biz.livekitService.CreateJoinToken(*meeting.Id, requester)
	if err != nil {
		return "", common.ErrInternal(errors.Wrap(err, "can not create join token"))

	}
	return token, nil
}
