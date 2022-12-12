package ginre

import (
	"bds-service/common"
	"bds-service/component"
	"bds-service/modules/realestate/rebiz"
	"bds-service/modules/realestate/restore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchByKeyword(context component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		keyword := c.Query("keyword")

		store := restore.NewSQLStore(context.GetMainDbConnection())
		biz := rebiz.NewSearchByKeywordBiz(store)
		result, err := biz.SearchByKeyword(c.Request.Context(), keyword)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
