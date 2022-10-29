package component

import (
	"bds-service/component/tokenprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDbConnection() *gorm.DB
	GetJWTProvider() tokenprovider.Provider
	GetJWTRefreshProvider() tokenprovider.Provider
}

type appCtx struct {
	db                   *gorm.DB
	tokenProvider        tokenprovider.Provider
	refreshTokenProvider tokenprovider.Provider
}

func (ctx appCtx) GetMainDbConnection() *gorm.DB {
	return ctx.db
}

func (ctx appCtx) GetJWTProvider() tokenprovider.Provider {
	return ctx.tokenProvider
}
func (ctx appCtx) GetJWTRefreshProvider() tokenprovider.Provider {
	return ctx.refreshTokenProvider
}

func NewAppContext(
	db *gorm.DB,
	provider tokenprovider.Provider,
	refreshProvider tokenprovider.Provider,
) *appCtx {
	return &appCtx{
		db:                   db,
		tokenProvider:        provider,
		refreshTokenProvider: refreshProvider,
	} //uploadProvider: uploadProvider

}
