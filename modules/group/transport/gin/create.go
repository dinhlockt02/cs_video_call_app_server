package groupgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	groupbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/group/biz"
	groupmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/group/model"
	groupstore "github.com/dinhlockt02/cs_video_call_app_server/modules/group/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGroup(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data *groupmdl.Group

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		createGroupBiz := groupbiz.NewCreateGroupBiz(groupStore)

		if err := createGroupBiz.Create(c.Request.Context(), requester.GetId(), data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, gin.H{"data": data.Id})
	}
}
