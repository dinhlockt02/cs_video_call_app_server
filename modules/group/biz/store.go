package groupbiz

import (
	"context"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
)

type GroupStore interface {
	Create(ctx context.Context, group *groupmdl.Group) error
}
