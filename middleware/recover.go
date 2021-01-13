package middleware

import (
	"errors"
	"fmt"
	"fooddlv/appctx"
	"fooddlv/common"
	"github.com/gin-gonic/gin"
)

func Recover(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
					return
				}

				appErr := common.ErrInternal(errors.New(fmt.Sprintf("%v", err)))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
