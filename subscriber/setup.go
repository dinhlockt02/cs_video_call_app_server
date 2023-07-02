package subscriber

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
)

func Setup(ctx context.Context, appCtx appcontext.AppContext) {
	UpdateMeetingStateWhenRoomFinished(ctx, appCtx)
}
