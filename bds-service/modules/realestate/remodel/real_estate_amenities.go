package remodel

type RealEstateAmenity struct {
	ReId      int `json:"re_id" gorm:"column:re_id;primaryKey"`
	AmenityId int `json:"amenity_id" gorm:"column:amenity_id;primaryKey"`
}

func (r *RealEstateAmenity) TableName() string {
	return "real_estate_amenities"
}
