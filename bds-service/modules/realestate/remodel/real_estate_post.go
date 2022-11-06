package remodel

import (
	"bds-service/common"
	"time"
)

type RealEstatePost struct {
	common.SQLModel
	ReId int `json:"-" gorm:"column:re_id;"`
	//belongs to
	RealEstate RealEstate `json:"real_estate" gorm:"foreignKey:re_id;references:id"`
	//belongs to
	RePostType  RealEstatePostType `json:"re_post_type" gorm:"foreignKey:post_type_id;references:id"`
	PostTypeId  int                `json:"-" gorm:"column:post_type_id;"`
	StartDate   time.Time          `json:"start_date" gorm:"column:start_date;"`
	Duration    int64              `json:"duration" gorm:"column:duration;"`
	AutoRenew   bool               `json:"auto_renew" gorm:"column:auto_renew;"`
	Title       string             `json:"title" gorm:"column:title;"`
	Description string             `json:"description" gorm:"column:description;"`
}

func (p RealEstatePost) TableName() string {
	return "real_estate_posts"
}
