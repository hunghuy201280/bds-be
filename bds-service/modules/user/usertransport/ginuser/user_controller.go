package ginuser

import (
	"bds-service/component"
	"bds-service/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(engine *gin.Engine, ctx component.AppContext) {
	users := engine.Group("/users")

	users.POST("", Register(ctx))
	users.GET("/:id", middleware.Authenticate(ctx), GetUserInfo(ctx))
	users.POST("/login", Login(ctx))
	users.POST("/refresh_token", RefreshToken(ctx))
}
