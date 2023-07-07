package subscriber

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
)

func Setup(ctx context.Context, appCtx appcontext.AppContext) {
	UpdateMeetingOrCallStateWhenRoomFinished(ctx, appCtx)
	UpdateGroupWhenRoomFinished(ctx, appCtx)
	UpdateGroupWhenRoomCreated(ctx, appCtx)
	UpdateCallsWhenUserUpdateProfile(ctx, appCtx)
	UpdateMessagesWhenUserUpdateProfile(ctx, appCtx)
	UpdateMeetingParticipantsWhenUserUpdateProfile(ctx, appCtx)
	UpdateRequestsWhenUserUpdateProfile(ctx, appCtx)
	UpdateRequestsWhenGroupUpdated(ctx, appCtx)
	UpdateNotificationWhenUserUpdateProfile(ctx, appCtx)
	UpdateNotificationWhenGroupUpdated(ctx, appCtx)
	DeleteNotificationWhenRequestDeleted(ctx, appCtx)
}
