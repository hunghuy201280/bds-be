package entitycommon

import (
	"database/sql/driver"
)

type EntityStatus int64

const (
	NORMAL  EntityStatus = 1
	DELETED EntityStatus = 2
)

func (e *EntityStatus) Scan(value any) error {

	*e = EntityStatus(value.(int64))
	return nil
}

func (e EntityStatus) Value() (driver.Value, error) {
	return int64(e), nil
}
