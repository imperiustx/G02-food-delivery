package ginuser

import (
	"fooddlv/appctx"
	"fooddlv/appctx/hasher"
	"fooddlv/common"
	"fooddlv/module/user/userbusiness"
	"fooddlv/module/user/usermodel"
	"fooddlv/module/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		repo := userbusiness.NewRegisterBusiness(store, md5)

		if err := repo.Register(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
