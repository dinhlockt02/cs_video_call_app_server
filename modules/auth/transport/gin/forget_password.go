package authgin

import (
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	authbiz "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/biz"
	authmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/model"
	authredis "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/redis"
	authstore "github.com/dinhlockt02/cs_video_call_app_server/modules/auth/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func ForgetPassword(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		type Body struct {
			Email string `json:"email"`
		}
		body := &Body{}

		err := c.ShouldBind(body)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if m := common.EmailRegexp.Match([]byte(body.Email)); !m {
			panic(common.ErrInvalidRequest(errors.New(authmodel.InvalidEmail)))

		}

		authStore := authstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		err = authbiz.
			NewForgetPasswordBiz(
				appCtx.Mailer(),
				authStore,
				authredis.NewRedisStore(
					appCtx.Redis(),
				),
			).
			Execute(c.Request.Context(), body.Email)
		if err != nil {
			panic(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
