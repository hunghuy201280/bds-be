package remodel

type RealEstateAmenity struct {
	ReId      int `json:"re_id" gorm:"column:re_id;"`
	AmenityId int `json:"amenity_id" gorm:"column:amenity_id;"`
}

func (r *RealEstateAmenity) TableName() string {
	return "real_estate_amenities"
}
