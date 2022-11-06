package ginre

import (
	"bds-service/component"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(engine *gin.Engine, ctx component.AppContext) {
	realEstates := engine.Group("/real-estates")

	realEstates.POST("", CreateRealEstate(ctx))
	realEstates.POST("/posts", CreateRealEstatePost(ctx))
	realEstates.POST("/posts/list", ListRealEstatePost(ctx))
	realEstates.POST("/insert-fake-data", CreateFakeData(ctx))
	realEstates.POST("/list", ListRealEstate(ctx))
	realEstates.GET("/app-profile", GetAppProfile(ctx))
}
