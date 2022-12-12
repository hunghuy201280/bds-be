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
	"strings"
)
import "github.com/jaswdr/faker"

var (
	ll = l.New()
)

func CreateRealEstate(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data remodel.RealEstateCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		ll.Info("create real estate body {}", l.Object("body", data))
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

		//id2 := 2
		//id3 := 3
		data := remodel.RealEstateCreate{
			ReTypeId:      common.RandInt(1, 20),
			ProvinceId:    "29",
			DistrictId:    "039",
			WardId:        "23042",
			Address:       fake.Address().Address(),
			Latitude:      fake.Float32(2, 1, 90),
			Longitude:     fake.Float32(2, 1, 90),
			Price:         fake.Float64(2, 9000000, 90000000),
			OwnerId:       common.RandInt(1, 20),
			Floors:        common.RandInt(1, 20),
			Area:          fake.Float32(2, 70, 2000),
			NoBedrooms:    common.RandInt(1, 20),
			NoWC:          common.RandInt(1, 20),
			HouseFacing:   remodel.NORTH,
			BalconyFacing: remodel.NORTH_EAST,
			Reason:        fake.Lorem().Sentence(30),
			BuiltAt:       strconv.Itoa(int(fake.Float32(0, 1990, 2020))),
			Documents:     strings.Join(strings.Split(fake.Lorem().Sentence(10), " "), ";"),
			Interiors:     strings.Join(strings.Split(fake.Lorem().Sentence(10), " "), ";"),
			Amenities: []remodel.Amenity{
				{
					Id: 3,
				},
				{
					Id: 4,
				},
				{
					Id: 5,
				},
			},
			Images: []common.Image{
				{
					Id: 4,
				},
				{
					Id: 5,
				},
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
