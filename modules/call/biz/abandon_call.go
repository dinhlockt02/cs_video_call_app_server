package callbiz

import (
	"context"
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	callrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/call/repository"
	callstore "github.com/dinhlockt02/cs_video_call_app_server/modules/call/store"
)

type abandonCallBiz struct {
	notification notirepo.NotificationRepository
	callRepo     callrepo.Repository
}

func NewAbandonCallBiz(
	notification notirepo.NotificationRepository,
	callRepo callrepo.Repository,
) *abandonCallBiz {
	return &abandonCallBiz{
		notification: notification,
		callRepo:     callRepo,
	}
}

// AbandonCall is the use case that user abandon call before other side answer.
func (biz *abandonCallBiz) AbandonCall(ctx context.Context, callerId string, calleeId string, callroom string) error {

	filter := common.GetAndFilter(
		callstore.GetCalleeIdFilter(calleeId),
		callstore.GetCallerIdFilter(callerId),
		callstore.GetCallStatusFilter(callmdl.OnGoing),
		callstore.GetCallRoomFilter(callroom),
	)
	call, err := biz.callRepo.FindCall(
		ctx,
		filter,
	)
	if err != nil {
		return err
	}
	if call == nil {
		return common.ErrEntityNotFound("Call", errors.New("call not found"))
	}

	call.Status = callmdl.Missed
	err = biz.callRepo.UpdateCall(ctx, filter, call)
	if err != nil {
		return err
	}
	err = biz.notification.CreateRejectIncomingCallNotification(
		ctx,
		callerId,
		&notimodel.NotificationObject{
			Id: calleeId,
		},
		&notimodel.NotificationObject{
			Id: callerId,
		},
		&notimodel.NotificationObject{
			Id:   *call.Id,
			Name: callroom,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
