package notigin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	notistore "github.com/dinhlockt02/cs_video_call_app_server/components/notification/store"
	notibiz "github.com/dinhlockt02/cs_video_call_app_server/modules/notification/biz"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotificationSetting(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {

		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)
		requesterId := requester.GetId()

		store := notistore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		repo := notirepo.NewNotificationRepository(store)
		biz := notibiz.NewGetNotificationSettingBiz(repo)
		rs, err := biz.Get(context.Request.Context(), requesterId)

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{"data": rs})
	}
}
