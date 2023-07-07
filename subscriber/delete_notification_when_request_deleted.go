package subscriber

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notistore "github.com/dinhlockt02/cs_video_call_app_server/components/notification/store"
)

func DeleteNotificationWhenRequestDeleted(ctx context.Context, appCtx appcontext.AppContext) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicRequestDeleted)

	notificationStore := notistore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {

		for reqId := range ch {
			go func(rid string) {
				defer common.Recovery()
				err := notificationStore.Delete(ctx, notistore.GetDirectFilter(rid, notimodel.Request))
				if err != nil {
					panic(err)
				}
			}(reqId)
		}
	}()
}
