package devicebiz

import (
	"context"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	devicestore "github.com/dinhlockt02/cs_video_call_app_server/modules/device/store"
)

type getDevicesBiz struct {
	store devicestore.Store
}

func NewGetDevicesBiz(store devicestore.Store) *getDevicesBiz {
	return &getDevicesBiz{store: store}
}

func (biz *getDevicesBiz) Get(ctx context.Context, filter map[string]interface{}) ([]*devicemodel.GetDeviceDto, error) {
	return biz.store.Get(ctx, filter)
}
