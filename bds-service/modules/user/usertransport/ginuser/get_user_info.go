package ginuser

import (
	"bds-service/common"
	"bds-service/component"
	"bds-service/modules/user/userbiz"
	"bds-service/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserInfo(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))

		}

		db := ctx.GetMainDbConnection()
		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewGetUserInfoBiz(store)
		user, err := biz.GetUserInfo(c.Request.Context(), userId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(http.StatusOK, "success", user, nil, nil))

	}
}
