package ginre

import (
	"bds-go-auth-service/component"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(engine *gin.Engine, ctx component.AppContext) {
	realEstates := engine.Group("/real-estates")

	realEstates.POST("", CreateRealEstate(ctx))
}
