package ginuser

import (
	"fooddlv/appctx"
	"fooddlv/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Masker)
		data.Mask(true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
