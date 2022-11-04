package remodel

import (
	"bds-service/common"
)

type RealEstate struct {
	common.SQLModel
	RealEstateType RealEstateType      `json:"real_estate_type" gorm:"foreignKey:ReId;references:Id"`
	ProvinceId     int                 `json:"province_id" gorm:"column:province_id;"`
	DistrictId     int                 `json:"district_id" gorm:"column:district_id;"`
	WardId         int                 `json:"ward_id" gorm:"column:ward_id;"`
	Address        string              `json:"address" gorm:"column:address;"`
	Latitude       float32             `json:"latitude" gorm:"column:latitude;"`
	Longitude      float32             `json:"longitude" gorm:"column:longitude;"`
	Price          float64             `json:"price" gorm:"column:price;"`
	OwnerId        int                 `json:"owner_id" gorm:"column:owner_id;"`
	Floors         int                 `json:"floors" gorm:"column:floors;"`
	Area           float32             `json:"area" gorm:"column:area;"`
	NoBedrooms     int                 `json:"no_bedrooms" gorm:"column:no_bedrooms;"`
	NoWC           int                 `json:"no_wc" gorm:"column:no_wc;"`
	HouseFacing    Direction           `json:"house_facing" gorm:"column:house_facing;"`
	BalconyFacing  Direction           `json:"balcony_facing" gorm:"column:balcony_facing;"`
	Reason         string              `json:"reason" gorm:"column:reason;"`
	BuiltAt        string              `json:"built_at" gorm:"column:built_at;"`
	Documents      string              `json:"documents" gorm:"column:documents;"`
	Amenities      []RealEstateAmenity `json:"amenities" gorm:"foreignKey:ReId;references:Id;"`
	Images         []common.Image      `json:"images" gorm:"many2many:real_estate_images;foreignKey:Id;joinForeignKey:ReId;References:Id;joinReferences:ImageId"`
}

func (receiver RealEstate) TableName() string {
	return "real_estates"
}
