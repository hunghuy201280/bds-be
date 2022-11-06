package remodel

import (
	"bds-service/common"
	"github.com/go-playground/validator/v10"
	"time"
)

type RealEstatePostCreate struct {
	common.SQLModel
	ReId        int       `json:"re_id" gorm:"column:re_id;" validate:"required"`
	PostTypeId  int       `json:"post_type_id" gorm:"column:post_type_id;" validate:"required"`
	StartDate   time.Time `json:"start_date" gorm:"column:start_date;" validate:"required"`
	Duration    int64     `json:"duration" gorm:"column:duration;" validate:"required"`
	AutoRenew   bool      `json:"auto_renew" gorm:"column:auto_renew;"`
	Title       string    `json:"title" gorm:"column:title;" validate:"required"`
	Description string    `json:"description" gorm:"column:description;" validate:"required"`
}

func (p RealEstatePostCreate) TableName() string {
	return RealEstatePost{}.TableName()
}

func (p *RealEstatePostCreate) Validate() error {

	validate := validator.New()
	return validate.Struct(p)
}
