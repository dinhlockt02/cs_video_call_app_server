package friendbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	friendmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/model"
	friendrepo "github.com/dinhlockt02/cs_video_call_app_server/modules/friend/repository"
)

type recallRequestBiz struct {
	friendRepo friendrepo.Repository
}

func NewRecallRequestBiz(friendRepo friendrepo.Repository) *recallRequestBiz {
	return &recallRequestBiz{
		friendRepo: friendRepo,
	}
}

func (biz *recallRequestBiz) RecallRequest(ctx context.Context, senderId string, receiverId string) error {
	existedRequest, err := biz.friendRepo.FindRequest(ctx, senderId, receiverId)
	if err != nil {
		return err
	}
	if existedRequest == nil {
		return common.ErrInvalidRequest(friendmodel.ErrRequestNotFound)
	}
	filter := make(map[string]interface{})
	err = common.AddIdFilter(filter, *existedRequest.Id)
	if err != nil {
		return err
	}
	err = biz.friendRepo.DeleteRequest(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
