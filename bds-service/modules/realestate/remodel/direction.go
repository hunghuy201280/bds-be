package remodel

import (
	"bds-service/common"
	"database/sql/driver"
)

type Direction string

const (
	EAST       Direction = "east"
	NORTH      Direction = "north"
	WEST       Direction = "west"
	SOUTH      Direction = "south"
	SOUTH_EAST Direction = "south_east"
	SOUTH_WEST Direction = "south_west"
	NORTH_EAST Direction = "north_east"
	NORTH_WEST Direction = "north_west"
)

func (e *Direction) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return common.ErrCannotParseValue("Direction", "string")
	}
	*e = Direction(bytes)
	return nil
}

func (e Direction) Value() (driver.Value, error) {
	return string(e), nil
}
