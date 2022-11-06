package ginre

import (
	"bds-service/common"
	"bds-service/common/l"
	"bds-service/component"
	"bds-service/modules/realestate/rebiz"
	"bds-service/modules/realestate/remodel"
	"bds-service/modules/realestate/restore"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListRealEstatePost(context component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter remodel.Filter
		var paging common.Paging
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBindQuery(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		ll.Info("{}\n---\n{}", l.Object("paging", paging), l.Object("filter", filter))
		store := restore.NewSQLStore(context.GetMainDbConnection())
		biz := rebiz.NewListRealEstatePostBiz(store)
		result, err := biz.ListRealEstatePost(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}
		if lastIndex := len(result) - 1; lastIndex >= 0 {
			paging.NextCursor = strconv.Itoa(result[lastIndex].Id)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(http.StatusOK, "success", result, paging, filter))
	}
}
