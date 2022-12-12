package entitycommon

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type EntityStatus int64

const (
	NORMAL  EntityStatus = 1
	DELETED EntityStatus = 2
)

func (e *EntityStatus) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	*e = EntityStatus(bytes[0])
	return nil
}

func (e EntityStatus) Value() (driver.Value, error) {
	return int64(e), nil
}
