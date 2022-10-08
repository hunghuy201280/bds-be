package ginuser

import (
	"bds-go-auth-service/component"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(engine *gin.Engine, ctx component.AppContext) {
	users := engine.Group("/users")

	users.POST("/", Register(ctx))
	users.GET("/:id", GetUserInfo(ctx))
	users.POST("/login", Login(ctx))
}
