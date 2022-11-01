package remodel

type RealEstateType struct {
	TypeId int  `json:"type_id" gorm:"column:type_id"`
	ReId   int  `json:"re_id" gorm:"column:re_id;primaryKey"`
	IsRent bool `json:"is_rent" gorm:"column:is_rent;"`
}

func (t RealEstateType) TableName() string {
	return "real_estate_types"
}
