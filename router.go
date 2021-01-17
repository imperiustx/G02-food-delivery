package main

import (
	"fooddlv/appctx"
	"fooddlv/middleware"
	"fooddlv/module/note/notetransport/ginnote"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine, appCtx appctx.AppContext) {
	r.Use(middleware.Recover(appCtx))
	v1 := r.Group("/v1")

	notes := v1.Group("/notes")
	{
		notes.GET("", ginnote.ListNote(appCtx))
		notes.POST("", ginnote.CreateNote(appCtx))
		notes.GET("/:note-id", ginnote.GetNote(appCtx))
		notes.PUT("/:note-id", ginnote.UpdateNote(appCtx))
		notes.DELETE("/:note-id", ginnote.DeleteNote(appCtx))
	}

}

func setupAdminRouter(r *gin.Engine, appCtx appctx.AppContext) {
	r.Use(middleware.Recover(appCtx))
	//admin := r.Group("/v1/admin")
}
