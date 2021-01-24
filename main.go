package main

import (
	"fooddlv/appctx"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("DB_CONNECTION_STRING")
	secret := os.Getenv("SECRET_KEY")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	appCtx := appctx.NewAppContext(db, secret)

	r := gin.Default()

	setupRouter(r, appCtx)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
