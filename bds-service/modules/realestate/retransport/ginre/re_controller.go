package ginre

import (
	"bds-service/component"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(engine *gin.Engine, ctx component.AppContext) {
	realEstates := engine.Group("/real-estates")

	realEstates.POST("", CreateRealEstate(ctx))
	realEstates.POST("/insert-fake-data", CreateFakeData(ctx))
	realEstates.POST("/list", ListRealEstate(ctx))
}
