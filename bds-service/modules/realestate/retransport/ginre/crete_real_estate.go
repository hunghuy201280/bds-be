package ginre

import (
	"bds-go-auth-service/common"
	"bds-go-auth-service/component"
	"bds-go-auth-service/modules/realestate/rebiz"
	"bds-go-auth-service/modules/realestate/remodel"
	"bds-go-auth-service/modules/realestate/restore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRealEstate(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data remodel.RealEstateCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := ctx.GetMainDbConnection()
		store := restore.NewSQLStore(db)
		biz := rebiz.NewCreateRealEstateBiz(store)
		id, err := biz.CreateRealEstate(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusCreated, common.NewSuccessResponse(http.StatusCreated, "created",
			common.JS{
				"id": id,
			}, nil, nil))
	}
}
