package callbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	lksv "github.com/dinhlockt02/cs_video_call_app_server/components/livekit_service"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"
	callrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/call/repository"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"time"
)

type CreateCallBiz struct {
	callRepo       callrepo.Repository
	livekitService lksv.LiveKitService
	notification   notirepo.INotificationService
}

func NewCreateCallBiz(
	callRepo callrepo.Repository,
	livekitService lksv.LiveKitService,
	notification notirepo.INotificationService,
) *CreateCallBiz {
	return &CreateCallBiz{
		callRepo:       callRepo,
		livekitService: livekitService,
		notification:   notification,
	}
}

func (biz *CreateCallBiz) Create(ctx context.Context,
	requesterId string, friendId string) (string, string, error) {
	log.Debug().Str("requesterId", requesterId).Any("friendId", friendId).Msg("create a call")

	// can not call self
	if requesterId == friendId {
		return "", "", common.ErrInvalidRequest(errors.New("can not call self"))
	}

	// Find requester
	requesterFilter, err := common.GetIdFilter(requesterId)
	if err != nil {
		return "", "", common.ErrInvalidRequest(errors.Wrap(err, "invalid requester id"))
	}

	requester, err := biz.callRepo.FindUser(ctx, requesterFilter)
	if err != nil {
		return "", "", common.ErrInternal(errors.Wrap(err, "can not find requester"))
	}

	if requester == nil {
		return "", "", common.ErrEntityNotFound(common.UserEntity, errors.Wrap(err, "requester not found"))
	}

	// Find friend
	friendFilter, err := common.GetIdFilter(friendId)
	if err != nil {
		return "", "", common.ErrInvalidRequest(errors.Wrap(err, "invalid friend id"))
	}

	friend, err := biz.callRepo.FindUser(ctx, friendFilter)
	if err != nil {
		return "", "", common.ErrInternal(errors.Wrap(err, "can not find friend"))
	}

	if friend == nil {
		return "", "", common.ErrEntityNotFound(common.UserEntity, errors.New("friend not found"))
	}

	// Create meeting
	call := &callmdl.Call{
		Caller: &callmdl.User{
			Id:     requesterId,
			Name:   requester.Name,
			Avatar: requester.Avatar,
		},
		Callee: &callmdl.User{
			Id:     friendId,
			Name:   friend.Name,
			Avatar: friend.Avatar,
		},
		Status:   callmdl.OnGoing,
		CalledAt: common.Ptr(time.Now()),
	}
	err = biz.callRepo.CreateCall(ctx, call)
	if err != nil {
		return "", "", common.ErrInternal(errors.Wrap(err, "can not create call"))
	}
	_, err = biz.livekitService.CreateRoom(ctx, *call.Id)
	if err != nil {
		return "", "", common.ErrInternal(errors.Wrap(err, "can not create livekit room"))
	}

	token, err := biz.livekitService.CreateJoinToken(*call.Id, requesterId)
	if err != nil {
		return "", "", common.ErrInternal(errors.Wrap(err, "can not create join token"))
	}

	// Notify friend
	go func() {
		err = biz.notification.CreateIncomingCallNotification(
			context.Background(),
			friend.Id,
			&notimodel.NotificationObject{
				Id:    requester.Id,
				Name:  requester.Name,
				Image: &requester.Avatar,
				Type:  notimodel.User,
			},
			&notimodel.NotificationObject{
				Id:    friend.Id,
				Name:  friend.Name,
				Image: &friend.Avatar,
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
	return token, *call.Id, nil
}
