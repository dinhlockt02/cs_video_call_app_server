package callbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	lksv "github.com/dinhlockt02/cs_video_call_app_server/components/livekit_service"
	"github.com/dinhlockt02/cs_video_call_app_server/components/pubsub"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	callrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/call/repository"
	callstore "github.com/dinhlockt02/cs_video_call_app_server/modules/call/store"
	meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type JoinCallBiz struct {
	callRepo       callrepo.Repository
	livekitService lksv.LiveKitService
	ps             pubsub.PubSub
}

func NewJoinCallBiz(
	callRepo callrepo.Repository,
	livekitService lksv.LiveKitService,
	ps pubsub.PubSub,
) *JoinCallBiz {
	return &JoinCallBiz{
		callRepo:       callRepo,
		livekitService: livekitService,
		ps:             ps,
	}
}

func (biz *JoinCallBiz) Join(ctx context.Context, requester, callId string) (string, error) {
	log.Debug().Str("requester", requester).
		Str("callId", callId).
		Msg("join call")

	callFilter, err := common.GetIdFilter(callId)
	if err != nil {
		return "", common.ErrInvalidRequest(errors.Wrap(err, "invalid call id"))
	}

	call, err := biz.callRepo.FindCall(ctx, common.GetAndFilter(callFilter, callstore.GetCallStatusFilter(callmdl.OnGoing)))
	if err != nil {
		return "", common.ErrInternal(errors.Wrap(err, "can not find call"))
	}
	if call == nil {
		return "", common.ErrEntityNotFound(common.CallEntity, errors.New(meetingmodel.MeetingNotFound))
	}

	if call.Callee.Id != requester {
		return "", common.ErrForbidden(errors.New("you not have permission to join call"))
	}

	token, err := biz.livekitService.CreateJoinToken(*call.Id, requester)
	if err != nil {
		return "", common.ErrInternal(errors.Wrap(err, "can not create join room token"))
	}
	err = biz.ps.Publish(ctx, common.TopicCallReacted, *call.Id)
	if err != nil {
		log.Error().Stack().Err(err).Msg("can not publish event")
	}
	return token, nil
}
