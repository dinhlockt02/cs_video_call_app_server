package callbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	lksv "github.com/dinhlockt02/cs_video_call_app_server/components/livekit_service"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	callrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/call/repository"
	callstore "github.com/dinhlockt02/cs_video_call_app_server/modules/call/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type RejectCallBiz struct {
	callRepo       callrepo.Repository
	livekitService lksv.LiveKitService
	notification   notirepo.INotificationService
}

func NewRejectCallBiz(
	callRepo callrepo.Repository,
	livekitService lksv.LiveKitService,
	notification notirepo.INotificationService,
) *RejectCallBiz {
	return &RejectCallBiz{
		callRepo:       callRepo,
		livekitService: livekitService,
		notification:   notification,
	}
}

func (biz *RejectCallBiz) Reject(ctx context.Context,
	requesterId string, callId string) error {
	log.Debug().Str("requesterId", requesterId).Str("callId", callId).Msg("abandon a call")

	// Find requester
	requesterFilter, err := common.GetIdFilter(requesterId)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid requester id"))
	}

	requester, err := biz.callRepo.FindUser(ctx, requesterFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find requester"))
	}

	if requester == nil {
		return common.ErrEntityNotFound(common.UserEntity, errors.Wrap(err, "requester not found"))
	}

	// Find call
	callFilter, err := common.GetIdFilter(callId)
	if err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "invalid call id"))
	}

	callFilter = common.GetAndFilter(callFilter, callstore.GetCallStatusFilter(callmdl.OnGoing))

	call, err := biz.callRepo.FindCall(ctx, callFilter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not find friend"))
	}

	if call == nil {
		return common.ErrEntityNotFound(common.CallEntity, errors.New("call not found"))
	}

	if call.Callee.Id != requesterId {
		return common.ErrForbidden(errors.New("you not have permission to abandon call"))
	}

	err = biz.callRepo.UpdateCall(ctx, callFilter, &callmdl.UpdateCall{
		Status: callmdl.Reject,
	})
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not update call"))
	}
	err = biz.livekitService.CloseRoom(ctx, *call.Id)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not close livekit room"))
	}
	// Notify friend
	go func() {
		err = biz.notification.CreateRejectIncomingCallNotification(
			context.Background(),
			call.Caller.Id,
			&notimodel.NotificationObject{
				Id:    requester.Id,
				Name:  requester.Name,
				Image: &requester.Avatar,
				Type:  notimodel.User,
			},
			&notimodel.NotificationObject{
				Id:    call.Caller.Id,
				Name:  call.Caller.Name,
				Image: &call.Caller.Avatar,
				Type:  notimodel.User,
			},
			&notimodel.NotificationObject{
				Id:    *call.Id,
				Name:  requester.Name,
				Image: &requester.Avatar,
				Type:  notimodel.CallRoom,
			},
		)
	}()
	return nil
}
