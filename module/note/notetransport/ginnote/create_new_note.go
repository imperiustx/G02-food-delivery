package ginnote

import (
	"fooddlv/appctx"
	"fooddlv/module/note/notemodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateNote(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetDBConnection()

		var data notemodel.Note
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&data)

		c.JSON(http.StatusOK, data)
	}
}
