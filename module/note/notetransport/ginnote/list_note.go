package ginnote

import (
	"fooddlv/appctx"
	"fooddlv/common"
	"fooddlv/module/note/notebusiness"
	"fooddlv/module/note/notemodel"
	"fooddlv/module/note/notestorage"
	"github.com/gin-gonic/gin"
)

func ListNote(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		var filter notemodel.Filter

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := appCtx.GetDBConnection()

		store := notestorage.NewSQLStore(db)
		//notes, err := store.ListNote()

		bizNote := notebusiness.NewListNoteBiz(store)
		notes, err := bizNote.ListAllNote(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSuccessResponse(notes, paging, filter))
	}
}
