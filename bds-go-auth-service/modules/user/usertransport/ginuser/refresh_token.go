package ginuser

import (
	"bds-go-auth-service/common"
	"bds-go-auth-service/component"
	"bds-go-auth-service/modules/user/userbiz"
	"bds-go-auth-service/modules/user/usermodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshToken(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserRefreshToken
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		biz := userbiz.NewRefreshTokenBiz(ctx.GetJWTRefreshProvider(), ctx.GetJWTProvider())
		token, refreshToken, err := biz.RefreshToken(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(http.StatusOK, "success", common.JS{
			"token":         token,
			"refresh_token": refreshToken,
		}, nil, nil))

	}
}
