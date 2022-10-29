package remodel

import (
	"bds-service/common"
	"errors"
	"github.com/go-playground/validator/v10"
	"time"
)

type RealEstateCreate struct {
	common.SQLModel
	Address          string              `json:"address" gorm:"column:address;" validate:"max=200,required"`
	Latitude         float32             `json:"latitude" gorm:"column:latitude;" validate:"required,gte=-90,lte=90" `
	Longitude        float32             `json:"longitude" gorm:"column:longitude;" validate:"required,gte=-180,lte=180" `
	MinPrice         float64             `json:"min_price" gorm:"column:min_price;" validate:"required,gte=0" `
	MaxPrice         float64             `json:"max_price" gorm:"column:max_price;" validate:"required,gte=0" `
	OwnerId          int                 `json:"owner_id" gorm:"column:owner_id;" validate:"required,gte=0" `
	Floors           int                 `json:"floors" gorm:"column:floors;" validate:"required,gte=0" `
	Area             float32             `json:"area" gorm:"column:area;" validate:"required,gte=0" `
	NoBedrooms       int                 `json:"no_bedrooms" gorm:"column:no_bedrooms;" validate:"required,gte=0" `
	NoWC             int                 `json:"no_wc" gorm:"column:no_wc;" validate:"required,gte=0" `
	ExpectedSoldDate time.Time           `json:"expected_sold_date" gorm:"column:expected_sold_date;" validate:"required"`
	Reason           string              `json:"reason" gorm:"column:reason;" `
	Description      string              `json:"description" gorm:"column:description;" `
	BuiltAt          string              `json:"built_at" gorm:"column:built_at;" validate:"required,len=4" `
	RealEstateTypeId int                 `json:"real_estate_type_id" gorm:"column:real_estate_type_id;" validate:"required" `
	ProvinceId       int                 `json:"province_id" gorm:"column:province_id;" validate:"required,gte=0" `
	DistrictId       int                 `json:"district_id" gorm:"column:district_id;" validate:"required,gte=0" `
	WardId           int                 `json:"ward_id" gorm:"column:ward_id;" validate:"required,gte=0" `
	Documents        string              `json:"documents" gorm:"column:documents;" validate:"required" `
	Amenities        []RealEstateAmenity `json:"amenities" gorm:"foreignKey:ReId"`
}

func (c RealEstateCreate) TableName() string {
	return RealEstate{}.TableName()
}

func (c *RealEstateCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}
	if c.MinPrice >= c.MaxPrice {
		return errors.New("min price should be less than max price")
	}
	return err

}
