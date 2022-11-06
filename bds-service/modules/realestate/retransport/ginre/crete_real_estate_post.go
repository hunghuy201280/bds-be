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
)

func CreateRealEstatePost(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data remodel.RealEstatePostCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		ll.Info("create real estate post body {}", l.Object("body", data))

		db := ctx.GetMainDbConnection()
		store := restore.NewSQLStore(db)
		biz := rebiz.NewCreatePostBiz(store)
		id, err := biz.CreatePost(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.JSON(
			http.StatusCreated,
			common.NewSuccessResponse(
				http.StatusCreated, "created",
				common.JS{
					"id": id,
				}, nil, nil,
			),
		)
	}
}
