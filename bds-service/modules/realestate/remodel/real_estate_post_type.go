package remodel

type RealEstatePostType struct {
	Id          int     `json:"id" gorm:"column:id;primaryKey"`
	Name        string  `json:"name" gorm:"column:name;"`
	PricePerDay float64 `json:"price_per_day" gorm:"column:price_per_day;"`
}

func (p RealEstatePostType) TableName() string {
	return "real_estate_post_types"
}
