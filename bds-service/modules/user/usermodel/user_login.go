package usermodel

import "github.com/go-playground/validator/v10"

type UserLogin struct {
	Phone    string `json:"phone" gorm:"column:phone;" validate:"required,numeric"`
	Password string `json:"password" gorm:"column:password;" validate:"required"`
}

func (u UserLogin) TableName() string {
	return "users"
}

func (u *UserLogin) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
