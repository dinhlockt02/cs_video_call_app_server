package authbiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/pkg/errors"
)

type LogoutDeviceStore interface {
	Delete(ctx context.Context, filter map[string]interface{}) error
}

type logoutBiz struct {
	deviceStore LogoutDeviceStore
}

func NewLogoutBiz(
	deviceStore LogoutDeviceStore,
) *logoutBiz {
	return &logoutBiz{
		deviceStore: deviceStore,
	}
}

func (biz *logoutBiz) Logout(ctx context.Context, filter map[string]interface{}) error {
	err := biz.deviceStore.Delete(ctx, filter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not delete user"))
	}
	return nil
}
