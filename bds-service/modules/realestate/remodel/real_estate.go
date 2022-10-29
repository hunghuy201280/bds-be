package remodel

import (
	"bds-service/common"
	"time"
)

type RealEstate struct {
	common.SQLModel
	Address          string              `json:"address" gorm:"column:address;"`
	Latitude         float32             `json:"latitude" gorm:"column:latitude;"`
	Longitude        float32             `json:"longitude" gorm:"column:longitude;"`
	MinPrice         float64             `json:"min_price" gorm:"column:min_price;"`
	MaxPrice         float64             `json:"max_price" gorm:"column:max_price;"`
	OwnerId          int                 `json:"owner_id" gorm:"column:owner_id;"`
	Floors           int                 `json:"floors" gorm:"column:floors;"`
	Area             float32             `json:"area" gorm:"column:area;"`
	NoBedrooms       int                 `json:"no_bedrooms" gorm:"column:no_bedrooms;"`
	NoWC             int                 `json:"no_wc" gorm:"column:no_wc;"`
	ExpectedSoldDate time.Time           `json:"expected_sold_date" gorm:"column:expected_sold_date;"`
	Reason           string              `json:"reason" gorm:"column:reason;"`
	Description      string              `json:"description" gorm:"column:description;"`
	BuiltAt          string              `json:"built_at" gorm:"column:built_at;"`
	RealEstateTypeId int                 `json:"real_estate_type_id" gorm:"column:real_estate_type_id;"`
	ProvinceId       int                 `json:"province_id" gorm:"column:province_id;"`
	DistrictId       int                 `json:"district_id" gorm:"column:district_id;"`
	WardId           int                 `json:"ward_id" gorm:"column:ward_id;"`
	Documents        string              `json:"documents" gorm:"column:documents;"`
	Amenities        []RealEstateAmenity `json:"amenities" gorm:"foreignKey:ReId"`
}

func (receiver RealEstate) TableName() string {
	return "real_estates"
}
