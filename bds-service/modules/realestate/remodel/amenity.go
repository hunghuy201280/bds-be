package remodel

type Amenity struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
}

func (a Amenity) TableName() string {
	return "amenities"
}
