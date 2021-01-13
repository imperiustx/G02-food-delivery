package main

import (
	"fooddlv/appctx"
	"fooddlv/middleware"
	"fooddlv/module/note/notemodel"
	"fooddlv/module/note/notetransport/ginnote"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Login struct {
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Note struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
	Status  int    `json:"status"  gorm:"column:status;"`
}

func (Note) TableName() string {
	return "notes"
}

type fakeListNoteStore struct{}

func (fakeListNoteStore) ListNote() ([]notemodel.Note, error) {
	return []notemodel.Note{
		{
			Id:    1,
			Title: "I am here",
		},
	}, nil
}

func main() {
	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	appCtx := appctx.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	//
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"user": Login{
	//			User:     "viettranx",
	//			Password: "123456",
	//		},
	//	})
	//})
	//
	v1 := r.Group("/v1")

	notes := v1.Group("/notes")
	{
		notes.GET("", ginnote.ListNote(appCtx))
		notes.POST("", ginnote.CreateNote(appCtx))
		notes.GET("/:note-id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": c.Param("note-id"),
			})
		})
		notes.PUT("/:note-id")
		notes.DELETE("/:note-id", ginnote.DeleteNote(appCtx))
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
