package meetingbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	grouprepo "github.com/dinhlockt02/cs_video_call_app_server/modules/group/repository"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	meetingrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/repository"
	meetingstore "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type GetMeetingBiz struct {
	meetingRepo meetingrepo.Repository
	groupRepo   grouprepo.Repository
}

func NewGetMeetingBiz(
	meetingRepo meetingrepo.Repository,
	groupRepo grouprepo.Repository,
) *GetMeetingBiz {
	return &GetMeetingBiz{
		meetingRepo, groupRepo,
	}
}

func (biz *GetMeetingBiz) Get(ctx context.Context, requesterId, groupId, meetingId string) (*meetingmodel.Meeting, error) {
	log.Debug().Str("requesterId", requesterId).
		Str("groupId", groupId).
		Str("meetingId", meetingId).
		Msg("join meeting")

	idFilter, err := common.GetIdFilter(meetingId)
	if err != nil {
		return nil, common.ErrInvalidRequest(errors.Wrap(err, "invalid meeting id"))
	}

	meeting, err := biz.meetingRepo.FindMeeting(ctx, common.GetAndFilter(idFilter, meetingstore.GetGroupFilter(groupId)))
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not find meeting"))
	}

	if meeting == nil {
		return nil, common.ErrEntityNotFound(common.MeetingEntity, errors.New("meeting not found"))
	}

	return meeting, nil
}
