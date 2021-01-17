package appctx

import (
	"gorm.io/gorm"
	"sync"
)

type AppContext interface {
	GetDBConnection() *gorm.DB
}

type appContext struct {
	db       *gorm.DB
	o        *sync.Once
	isLoaded bool
}

func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{db: db, o: new(sync.Once)}
}

func (ctx *appContext) GetDBConnection() *gorm.DB {
	//ctx.o.Do(func() {
	//	time.Sleep(time.Second * 10)
	//})
	//
	//go func() {
	//	for {
	//		time.Sleep(time.Second * 5)
	//		ctx.db.DisableAutomaticPing
	//		if err := ctx.db.Ping(); err != nil {
	//			ctx.o = new(sync.Once)
	//		}
	//	}
	//}()

	return ctx.db.Session(&gorm.Session{NewDB: true})
}
