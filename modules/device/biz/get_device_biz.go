package devicebiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	devicestore "github.com/dinhlockt02/cs_video_call_app_server/modules/device/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type GetDevicesBiz struct {
	store devicestore.Store
}

func NewGetDevicesBiz(store devicestore.Store) *GetDevicesBiz {
	return &GetDevicesBiz{store: store}
}

func (biz *GetDevicesBiz) Get(ctx context.Context, filter map[string]interface{}) ([]*devicemodel.GetDeviceDto, error) {
	log.Debug().Any("filter", filter).Msg("get devices")
	devices, err := biz.store.Get(ctx, filter)
	if err != nil {
		return nil, common.ErrInternal(errors.Wrap(err, "can not get devices"))
	}
	return devices, nil
}
