package ginnote

import (
	"fooddlv/appctx"
	"fooddlv/common"
	"fooddlv/module/note/notebusiness"
	"fooddlv/module/note/notemodel"
	"fooddlv/module/note/notestorage"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UpdateNote(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data notemodel.NoteUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()

		idString := c.Param("note-id")
		id, _ := strconv.Atoi(idString)

		store := notestorage.NewSQLStore(db)
		biz := notebusiness.NewUpdateNoteBiz(store)

		err := biz.UpdateNote(c.Request.Context(), id, &data)

		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
