package common

import (
	"bds-go-auth-service/common/entitycommon"
	"time"
)

type SQLModel struct {
	Id        int                       `json:"id" gorm:"column:id;"`
	Status    entitycommon.EntityStatus `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time                `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time                `json:"updated_at" gorm:"column:updated_at;"`
}
