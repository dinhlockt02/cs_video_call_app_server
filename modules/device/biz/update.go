package devicebiz

import (
	"context"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	devicestore "github.com/dinhlockt02/cs_video_call_app_server/modules/device/store"
)

type createDeviceBiz struct {
	store devicestore.Store
}

func NewUpdateDeviceBiz(store devicestore.Store) *createDeviceBiz {
	return &createDeviceBiz{store: store}
}

func (biz *createDeviceBiz) Update(ctx context.Context, filter map[string]interface{}, data *devicemodel.UpdateDevice) error {

	if err := data.Process(); err != nil {
		return err
	}

	err := biz.store.Update(ctx, filter, data)
	if err != nil {
		return err
	}

	return nil
}
