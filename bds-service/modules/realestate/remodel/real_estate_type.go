package remodel

type RealEstateType struct {
	Id     int    `json:"id" gorm:"column:id;primaryKey"`
	Name   string `json:"name" gorm:"column:name"`
	IsRent bool   `json:"is_rent" gorm:"column:is_rent"`
}

func (t RealEstateType) TableName() string {
	return "real_estate_types"
}
