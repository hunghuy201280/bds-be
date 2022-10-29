package ginre

import (
	"bds-service/common"
	"bds-service/common/entitycommon"
	"bds-service/component"
	"bds-service/modules/realestate/rebiz"
	"bds-service/modules/realestate/remodel"
	"bds-service/modules/realestate/restore"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)
import "github.com/jaswdr/faker"

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
func CreateFakeData(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fake := faker.New()

		data := remodel.RealEstateCreate{
			Address:          fake.Address().Address(),
			Latitude:         fake.Float32(2, 1, 90),
			Longitude:        fake.Float32(2, 1, 90),
			MinPrice:         fake.Float64(2, 1_000_000, 9_000_000),
			MaxPrice:         fake.Float64(2, 9_000_001, 10_000_000),
			OwnerId:          common.RandInt(1, 20),
			Floors:           common.RandInt(1, 20),
			Area:             fake.Float32(2, 70, 2000),
			NoBedrooms:       common.RandInt(1, 20),
			NoWC:             common.RandInt(1, 20),
			ExpectedSoldDate: fake.Time().TimeBetween(time.Now(), time.Date(2024, 1, 1, 1, 1, 1, 1, time.Local)),
			Description:      fake.Lorem().Sentence(30),
			Reason:           fake.Lorem().Sentence(30),
			BuiltAt:          strconv.Itoa(int(fake.Float32(0, 1990, 2020))),
			RealEstateTypeId: common.RandInt(1, 20),
			ProvinceId:       common.RandInt(1, 20),
			DistrictId:       common.RandInt(1, 20),
			WardId:           common.RandInt(1, 20),
			Documents:        strings.Join(strings.Split(fake.Lorem().Sentence(10), " "), ";"),
			SQLModel: common.SQLModel{
				Status: entitycommon.NORMAL,
			},
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
