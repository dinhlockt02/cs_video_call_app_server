package callbiz

import (
	"context"
	"errors"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	callrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/call/repository"
	"github.com/rs/zerolog/log"
	"time"
)

type createNewCallBiz struct {
	notification notirepo.NotificationRepository
	callRepo     callrepo.Repository
}

func NewCreateNewCallBiz(
	notification notirepo.NotificationRepository,
	callRepo callrepo.Repository,
) *createNewCallBiz {
	return &createNewCallBiz{
		notification: notification,
		callRepo:     callRepo,
	}
}

// CreateNewCall create a new call with status of waiting.
func (biz *createNewCallBiz) CreateNewCall(ctx context.Context, callerId string, calleeId string, callroom string) error {

	// Find caller
	filter := make(map[string]interface{})
	_ = common.AddIdFilter(filter, callerId)
	caller, err := biz.callRepo.FindUser(ctx, filter)
	if err != nil {
		return err
	}
	if caller == nil {
		return common.ErrEntityNotFound("User", errors.New("caller not found"))
	}

	// Find Callee
	filter = make(map[string]interface{})
	_ = common.AddIdFilter(filter, calleeId)
	callee, err := biz.callRepo.FindUser(ctx, filter)
	if callee == nil {
		return common.ErrEntityNotFound("User", errors.New("callee not found"))
	}

	now := time.Now()
	// Create new call

	call := callmdl.Call{
		Caller:   caller,
		Callee:   callee,
		Status:   callmdl.OnGoing,
		CalledAt: &now,
		CallRoom: callroom,
	}

	err = biz.callRepo.CreateCall(ctx, &call)
	if err != nil {
		return err
	}
	err = biz.notification.CreateIncomingCallNotification(
		ctx,
		callee.Id,
		&notimodel.NotificationObject{
			Id:    caller.Id,
			Name:  caller.Name,
			Image: &caller.Avatar,
			Type:  notimodel.User,
		},
		&notimodel.NotificationObject{
			Id:    callee.Id,
			Name:  callee.Name,
			Image: &caller.Avatar,
			Type:  notimodel.User,
		},
		&notimodel.NotificationObject{
			Id:    *call.Id,
			Name:  callroom,
			Image: nil,
			Type:  notimodel.CallRoom,
		},
	)

	log.Debug().Msgf("%v", call)
	if err != nil {
		return err
	}

	return nil
}
