package devicebiz

import (
	"context"
	devicestore "github.com/dinhlockt02/cs_video_call_app_server/modules/device/store"
)

type deleteDeviceBiz struct {
	store devicestore.Store
}

func NewDeleteDevicesBiz(store devicestore.Store) *deleteDeviceBiz {
	return &deleteDeviceBiz{store: store}
}

func (biz *deleteDeviceBiz) Delete(ctx context.Context, filter map[string]interface{}) error {
	return biz.store.Delete(ctx, filter)
}
