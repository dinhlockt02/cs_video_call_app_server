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

type rejectCallBiz struct {
	notification notirepo.NotificationRepository
	callRepo     callrepo.Repository
}

func NewRejectCallBiz(
	notification notirepo.NotificationRepository,
	callRepo callrepo.Repository,
) *rejectCallBiz {
	return &rejectCallBiz{
		notification: notification,
		callRepo:     callRepo,
	}
}

// Reject is the use case that user reject incoming call.
func (biz *rejectCallBiz) Reject(ctx context.Context, callerId string, calleeId string, callroom string) error {

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

	call.Status = callmdl.Reject
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
