package devicebiz

import (
	"context"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
)

type DeviceStore interface {
	Create(ctx context.Context, data *devicemodel.Device) error
	Update(ctx context.Context, filter map[string]interface{}, data *devicemodel.UpdateDevice) error
}
