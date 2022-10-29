package remodel

type Filter struct {
	Id                       int     `json:"id" form:"id"`
	MinPrice                 float64 `json:"min_price" form:"min_price"`
	MaxPrice                 float64 `json:"max_price" form:"max_price"`
	NoBedrooms               int     `json:"no_bedrooms" form:"no_bedrooms"`
	NoWC                     int     `json:"no_wc" form:"no_wc"`
	MinArea                  float32 `json:"min_area" form:"area"`
	MaxArea                  float32 `json:"max_area" form:"area"`
	Floors                   int     `json:"floors" gorm:"floors"`
	RealEstateTypeIds        []int   `json:"real_estate_type_ids" form:"real_estate_type_ids"`
	RealEstateAmenityTypeIds []int   `json:"real_estate_amenity_type_ids" form:"real_estate_amenity_ids"`
	ProvinceId               int     `json:"province_id" form:"province_id"`
	DistrictId               int     `json:"district_id" form:"district_id"`
	WardId                   int     `json:"ward_id" form:"ward_id"`
}
