package remodel

import "bds-service/common"

type RealEstateImage struct {
	ReId    int          `json:"re_id" gorm:"column:re_id;primaryKey" validate:"required,lte=0"`
	ImageId int          `json:"id" gorm:"column:image_id;primaryKey" validate:"required,lte=0"`
	Image   common.Image `json:"image"`
}

func (r RealEstateImage) TableName() string {
	return "real_estate_images"
}
