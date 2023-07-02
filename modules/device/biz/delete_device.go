package devicebiz

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	devicestore "github.com/dinhlockt02/cs_video_call_app_server/modules/device/store"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type DeleteDeviceBiz struct {
	store devicestore.Store
}

func NewDeleteDevicesBiz(store devicestore.Store) *DeleteDeviceBiz {
	return &DeleteDeviceBiz{store: store}
}

func (biz *DeleteDeviceBiz) Delete(ctx context.Context, filter map[string]interface{}) error {
	log.Debug().Any("filter", filter).Msg("delete devices")
	err := biz.store.Delete(ctx, filter)
	if err != nil {
		return common.ErrInternal(errors.Wrap(err, "can not delete devices"))
	}
	return nil
}
