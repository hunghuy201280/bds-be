package ginre

import (
	"bds-service/common"
	"bds-service/component"
	"bds-service/modules/realestate/rebiz"
	"bds-service/modules/realestate/restore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAppProfile(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDbConnection()
		store := restore.NewSQLStore(db)
		biz := rebiz.NewGetAppProfileBiz(store)
		appProfile, err := biz.GetAppProfile(c.Request.Context())
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(appProfile))
	}
}
