package remodel

import (
	"bds-service/common"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/goutil"
)

type RealEstateCreate struct {
	common.SQLModel
	RealEstateType RealEstateType      `json:"real_estate_type" gorm:"foreignKey:ReId;references:Id" validate:"required"`
	ProvinceId     int                 `json:"province_id" gorm:"column:province_id;" validate:"required,gte=0"`
	DistrictId     int                 `json:"district_id" gorm:"column:district_id;" validate:"required,gte=0"`
	WardId         int                 `json:"ward_id" gorm:"column:ward_id;" validate:"required,gte=0"`
	Address        string              `json:"address" gorm:"column:address;" validate:"max=255,required"`
	Latitude       float32             `json:"latitude" gorm:"column:latitude;" validate:"required,gte=-90,lte=90"`
	Longitude      float32             `json:"longitude" gorm:"column:longitude;" validate:"required,gte=-180,lte=180"`
	Price          float64             `json:"price" gorm:"column:price;" validate:"required,gte=0"`
	OwnerId        int                 `json:"owner_id" gorm:"column:owner_id;" validate:"required,gte=0"`
	Floors         int                 `json:"floors" gorm:"column:floors;" validate:"required,gte=0"`
	Area           float32             `json:"area" gorm:"column:area;" validate:"required,gte=0"`
	NoBedrooms     int                 `json:"no_bedrooms" gorm:"column:no_bedrooms;" validate:"required,gte=0"`
	NoWC           int                 `json:"no_wc" gorm:"column:no_wc;" validate:"required,gte=0"`
	HouseFacing    Direction           `json:"house_facing" gorm:"column:house_facing;"`
	BalconyFacing  Direction           `json:"balcony_facing" gorm:"column:balcony_facing;"`
	Reason         string              `json:"reason" gorm:"column:reason;"`
	BuiltAt        string              `json:"built_at" gorm:"column:built_at;" validate:"required,len=4"`
	Documents      string              `json:"documents" gorm:"column:documents;"`
	Interiors      string              `json:"interiors" gorm:"column:interiors;"`
	Amenities      []RealEstateAmenity `json:"amenities" gorm:"foreignKey:ReId;references:Id;association_foreignkey:Id" validate:"required"`
	Images         []common.Image      `json:"images" gorm:"many2many:real_estate_images;foreignKey:Id;joinForeignKey:ReId;References:Id;joinReferences:ImageId"`
}

func (c RealEstateCreate) TableName() string {
	return RealEstate{}.TableName()
}

func (c *RealEstateCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if !goutil.IsEmpty(c.Images) {
		for _, item := range c.Images {
			err = validate.Struct(&item)
		}
	}
	if err != nil {
		return err
	}
	return err

}
