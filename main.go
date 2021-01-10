package main

import (
	"fooddlv/module/note/notebusiness"
	"fooddlv/module/note/notemodel"
	"fooddlv/module/note/notestorage"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
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

	r := gin.Default()
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
		notes.GET("", func(c *gin.Context) {
			//store := notestorage.NewSQLStore(db)
			//notes, err := store.ListNote()

			bizNote := notebusiness.NewListNoteBiz(fakeListNoteStore{})
			notes, err := bizNote.ListAllNote()

			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, notes)
		})

		notes.POST("", func(c *gin.Context) {
			var data Note
			if err := c.ShouldBindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			db.Create(&data)

			c.JSON(http.StatusOK, data)
		})
		notes.GET("/:note-id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": c.Param("note-id"),
			})
		})
		notes.PUT("/:note-id")
		notes.DELETE("/:note-id", func(c *gin.Context) {
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
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
