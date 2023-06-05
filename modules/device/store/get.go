package devicestore

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	devicemodel "github.com/dinhlockt02/cs_video_call_app_server/modules/device/model"
)

func (s *mongoStore) Get(ctx context.Context, filter map[string]interface{}) ([]*devicemodel.GetDeviceDto, error) {
	var devices []*devicemodel.GetDeviceDto
	cursor, err := s.database.Collection(devicemodel.GetDeviceDto{}.CollectionName()).Find(ctx, filter)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	if err = cursor.All(ctx, &devices); err != nil {
		return nil, common.ErrInvalidRequest(err)
	}
	return devices, nil
}
