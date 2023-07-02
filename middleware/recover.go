package middleware

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Recover(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if err == context.Canceled {
					return
				}
				c.Header("Content-Type", "application/json")
				appErr, ok := err.(*common.AppError)
				if !ok {
					appErr = common.ErrInternal(err.(error))
				}
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				log.Error().Stack().Err(appErr.RootError()).Send()
				return
			}
		}()
		c.Next()
	}
}
