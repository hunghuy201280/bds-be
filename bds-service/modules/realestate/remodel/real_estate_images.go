package remodel

type RealEstateImage struct {
	ReId    *int `json:"re_id" gorm:"column:re_id;primaryKey" validate:"required"`
	ImageId *int `json:"id" gorm:"column:image_id;primaryKey" validate:"required"`
}

func (r *RealEstateImage) TableName() string {
	return "real_estate_images"
}
