package appctx

import (
	"gorm.io/gorm"
)

type AppContext interface {
	GetDBConnection() *gorm.DB
	SecretKey() string
}

type appContext struct {
	db     *gorm.DB
	secret string
}

func NewAppContext(db *gorm.DB, secret string) *appContext {
	return &appContext{
		db:     db,
		secret: secret,
	}
}

func (ctx *appContext) GetDBConnection() *gorm.DB {
	return ctx.db.Session(&gorm.Session{NewDB: true})
}

func (ctx *appContext) SecretKey() string {
	return ctx.secret
}

type tokenExpiry struct {
	atExp int
	rtExp int
}

func NewTokenConfig() tokenExpiry {
	return tokenExpiry{
		atExp: 60 * 60 * 24 * 7,
		rtExp: 60 * 60 * 24 * 7 * 2,
	}
}

func (tk tokenExpiry) GetAtExp() int {
	return tk.atExp
}

func (tk tokenExpiry) GetRtExp() int {
	return tk.rtExp
}
