package usermodel

import "bds-service/common"

type UserFind struct {
	common.SQLModel
	Phone      string   `json:"phone" gorm:"column:phone;"`
	FacebookId string   `json:"fb_id" gorm:"column:fb_id;"`
	GoogleId   string   `json:"gg_id" gorm:"column:gg_id;"`
	LastName   string   `json:"last_name" gorm:"column:last_name;"`
	FirstName  string   `json:"first_name" gorm:"column:first_name;"`
	Role       UserRole `json:"role" gorm:"column:role;type:ENUM('buyer','seller','agency','admin')"`
	AvatarUrl  string   `json:"avatar_url,omitempty" gorm:"column:avatar_url;"`
}

func (u UserFind) TableName() string {
	return "users"
}
