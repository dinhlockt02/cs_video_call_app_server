package subscriber

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
)

func Setup(appCtx appcontext.AppContext, ctx context.Context) {
	UpdateMeetingStateWhenRoomFinished(appCtx, ctx)
}
