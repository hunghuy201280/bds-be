package usermodel

import (
	"bds-service/common"
	"github.com/go-playground/validator/v10"
)

type UserCreate struct {
	common.SQLModel
	Phone     string `json:"phone" gorm:"column:phone;" validate:"number,required"`
	FirstName string `json:"first_name" gorm:"column:first_name;" validate:"alpha,required"`
	LastName  string `json:"last_name" gorm:"column:last_name;" validate:"alpha,required"`
	Password  string `json:"password" gorm:"column:password;" validate:"required"`
	Salt      string `json:"-" gorm:"column:salt;"`
	Role      string `json:"-" gorm:"column:role;"`
}

func (c UserCreate) TableName() string {
	return User{}.TableName()
}

func (c *UserCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}
	return nil
}
