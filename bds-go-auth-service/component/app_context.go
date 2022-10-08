package component

import (
	"bds-go-auth-service/component/tokenprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDbConnection() *gorm.DB
	GetJWTProvider() tokenprovider.Provider
}

type appCtx struct {
	db            *gorm.DB
	tokenProvider tokenprovider.Provider
}

func (ctx appCtx) GetMainDbConnection() *gorm.DB {
	return ctx.db
}

func (ctx appCtx) GetJWTProvider() tokenprovider.Provider {
	return ctx.tokenProvider
}

func NewAppContext(
	db *gorm.DB,
	provider tokenprovider.Provider,
) *appCtx {
	return &appCtx{db: db, tokenProvider: provider} //uploadProvider: uploadProvider

}
