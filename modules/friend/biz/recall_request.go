package friendbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
)

type RecallRequestFriendStore interface {
	DeleteRequest(ctx context.Context, requestId string) error
	FindRequest(ctx context.Context, userId string, otherId string) (*friendmodel.Request, error)
}

type recallRequestBiz struct {
	friendStore RecallRequestFriendStore
}

func NewRecallRequestBiz(friendStore RecallRequestFriendStore) *recallRequestBiz {
	return &recallRequestBiz{
		friendStore: friendStore,
	}
}

func (biz *recallRequestBiz) RecallRequest(ctx context.Context, senderId string, receiverId string) error {
	existedRequest, err := biz.friendStore.FindRequest(ctx, senderId, receiverId)
	if err != nil {
		return err
	}
	if existedRequest == nil {
		return common.ErrInvalidRequest(friendmodel.ErrRequestNotFound)
	}
	err = biz.friendStore.DeleteRequest(ctx, *existedRequest.Id)
	if err != nil {
		return err
	}
	return nil
}
