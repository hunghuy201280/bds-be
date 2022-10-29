package usermodel

import "bds-service/common"

type UserLite struct {
	common.SQLModel
	Phone     string   `json:"phone" gorm:"column:phone;"`
	LastName  string   `json:"last_name" gorm:"column:last_name;"`
	FirstName string   `json:"first_name" gorm:"column:first_name;"`
	Password  string   `json:"-" gorm:"column:password;"`
	Salt      string   `json:"-" gorm:"column:salt;"`
	Role      UserRole `json:"role" gorm:"column:role;type:ENUM('buyer','seller','agency','admin')"`
	AvatarUrl string   `json:"avatar_url,omitempty" gorm:"column:avatar_url;"`
}

func (u UserLite) TableName() string {
	return "users"
}
