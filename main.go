package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"user": Login{
				User:     "viettranx",
				Password: "123456",
			},
		})
	})

	v1 := r.Group("/v1")

	notes := v1.Group("/notes")
	{
		notes.GET("")
		notes.POST("", func(c *gin.Context) {
			var data Login
			if err := c.ShouldBindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"user": data,
			})
		})
		notes.GET("/:note-id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": c.Param("note-id"),
			})
		})
		notes.PUT("/:note-id")
		notes.DELETE("/:note-id")
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
