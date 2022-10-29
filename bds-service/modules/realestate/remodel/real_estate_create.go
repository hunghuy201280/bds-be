package remodel

import (
	"bds-go-auth-service/common"
	"errors"
	"github.com/go-playground/validator/v10"
	"time"
)

type RealEstateCreate struct {
	common.SQLModel
	Address          string    `json:"address" gorm:"column:address;" validate:"max=50,required"`
	Latitude         float32   `json:"latitude" gorm:"column:latitude;" validate:"required" `
	Longitude        float32   `json:"longitude" gorm:"column:longitude;" validate:"required" `
	MinPrice         float64   `json:"min_price" gorm:"column:min_price;" validate:"required" `
	MaxPrice         float64   `json:"max_price" gorm:"column:max_price;" validate:"required" `
	OwnerId          int       `json:"owner_id" gorm:"column:owner_id;" validate:"required" `
	Floors           int       `json:"floors" gorm:"column:floors;" validate:"required" `
	Area             float32   `json:"area" gorm:"column:area;" validate:"required" `
	NoBedrooms       int       `json:"no_bedrooms" gorm:"column:no_bedrooms;" validate:"required" `
	NoWC             int       `json:"no_wc" gorm:"column:no_wc;" validate:"required" `
	ExpectedSoldDate time.Time `json:"expected_sold_date" gorm:"column:expected_sold_date;" validate:"required"`
	Reason           string    `json:"reason" gorm:"column:reason;" `
	Description      string    `json:"description" gorm:"column:description;" `
	BuiltAt          string    `json:"built_at" gorm:"column:built_at;" validate:"required,len=4" `
	RealEstateTypeId int       `json:"real_estate_type_id" gorm:"column:real_estate_type_id;" validate:"required" `
	ProvinceId       int       `json:"province_id" gorm:"column:province_id;" validate:"required" `
	DistrictId       int       `json:"district_id" gorm:"column:district_id;" validate:"required" `
	WardId           int       `json:"ward_id" gorm:"column:ward_id;" validate:"required" `
	Documents        string    `json:"documents" gorm:"column:documents;" validate:"required" `
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
