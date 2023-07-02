package devicebiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
	devicestore "github.com/dinhlockt02/cs_video_call_app_server/modules/device/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type CreateDeviceBiz struct {
	store devicestore.Store
}

func NewUpdateDeviceBiz(store devicestore.Store) *CreateDeviceBiz {
	return &CreateDeviceBiz{store: store}
}

func (biz *CreateDeviceBiz) Update(ctx context.Context,
	filter map[string]interface{}, data *devicemodel.UpdateDevice) error {
	log.Debug().Any("filter", filter).Any("data", data).Msg("update a device")
	if err := data.Process(); err != nil {
		return common.ErrInvalidRequest(errors.Wrap(err, "validate data failed"))
	}

	err := biz.store.Update(ctx, filter, data)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not update device"))
	}

	return nil
}
