package subscriber

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notistore "github.com/dinhlockt02/cs_video_call_app_server/components/notification/store"
)

func DeleteNotificationWhenCallReacted(ctx context.Context, appCtx appcontext.AppContext) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicCallReacted)

	notificationStore := notistore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {
		for callId := range ch {
			go func(callId string) {
				defer common.Recovery()
				err := notificationStore.Delete(ctx, notistore.GetPrepFilter(callId, notimodel.CallRoom))
				if err != nil {
					panic(err)
				}
			}(callId)
		}
	}()
}
