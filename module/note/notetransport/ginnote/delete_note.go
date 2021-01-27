package ginnote

import (
	"fooddlv/appctx"
	"fooddlv/common"
	"fooddlv/module/note/notebusiness"
	"fooddlv/module/note/notestorage"
	"github.com/gin-gonic/gin"
)

func DeleteNote(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetDBConnection()

		uid, err := common.FromBase58(c.Param("note-id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := notestorage.NewSQLStore(db)
		biz := notebusiness.NewDeleteNoteBiz(store, requester)

		err = biz.DeleteNote(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{"data": 1})
	}
}
