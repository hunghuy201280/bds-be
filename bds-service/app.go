package main

import (
	"bds-service/component"
	"bds-service/component/tokenprovider"
	"bds-service/component/uploadprovider"
	"bds-service/middleware"
	"bds-service/modules/realestate/retransport/ginre"
	"bds-service/modules/upload/uploadtransportation/ginupload"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RunService(
	db *gorm.DB,
	port string,
	provider tokenprovider.Provider,
	refreshProvider tokenprovider.Provider,
	uploadProvider uploadprovider.UploadProvider,
) error {

	appCtx := component.NewAppContext(db, provider, refreshProvider, uploadProvider)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	//r.Use(middleware.Authenticate(appCtx))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	},
	)

	ginre.RegisterHandler(r, appCtx)
	r.POST("/upload", ginupload.Upload(appCtx))

	return r.Run(":" + port)
}
