package notigin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	notimodel "github.com/dinhlockt02/cs_video_call_app_server/components/notification/model"
	notirepo "github.com/dinhlockt02/cs_video_call_app_server/components/notification/repository"
	notistore "github.com/dinhlockt02/cs_video_call_app_server/components/notification/store"
	notibiz "github.com/dinhlockt02/cs_video_call_app_server/modules/notification/biz"
	"github.com/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateNotificationSetting(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {

		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)
		requesterId := requester.GetId()

		var body notimodel.NotificationUser

		if err := context.ShouldBind(&body); err != nil {
			panic(common.ErrInvalidRequest(errors.Wrap(err, "invalid body data")))
		}

		store := notistore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		repo := notirepo.NewNotificationRepository(store)
		biz := notibiz.NewUpdateNotificationSettingBiz(repo)
		err := biz.Update(context.Request.Context(), requesterId, &body)

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
