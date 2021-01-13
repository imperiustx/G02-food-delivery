package ginnote

import (
	"fooddlv/appctx"
	"fooddlv/module/note/notebusiness"
	"fooddlv/module/note/notestorage"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteNote(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetDBConnection()

		idString := c.Param("note-id")
		id, _ := strconv.Atoi(idString)

		store := notestorage.NewSQLStore(db)
		biz := notebusiness.NewDeleteNoteBiz(store)

		err := biz.DeleteNote(id)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": 1})
	}
}
