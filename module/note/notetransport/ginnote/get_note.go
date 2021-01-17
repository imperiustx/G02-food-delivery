package ginnote

import (
	"fooddlv/appctx"
	"fooddlv/common"
	"fooddlv/module/note/notebusiness"
	"fooddlv/module/note/notestorage"
	"github.com/gin-gonic/gin"
)

func GetNote(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetDBConnection()

		uid, err := common.FromBase58(c.Param("note-id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := notestorage.NewSQLStore(db)

		bizNote := notebusiness.NewGetNoteBiz(store)
		data, err := bizNote.GetNote(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		data.GenUID(common.DBTypeNote, 1)

		c.JSON(200, common.SimpleSuccessResponse(data))
	}
}
