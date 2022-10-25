package usermodel

import "database/sql/driver"

type UserRole string

const (
	BUYER  UserRole = "buyer"
	SELLER UserRole = "seller"
	AGENCY UserRole = "agency"
	ADMIN  UserRole = "admin"
)

func (r UserRole) String() string {
	return string(r)
}

func (r *UserRole) Scan(value interface{}) error {
	*r = UserRole(value.([]byte))
	return nil
}

func (r UserRole) Value() (driver.Value, error) {
	return string(r), nil
}
