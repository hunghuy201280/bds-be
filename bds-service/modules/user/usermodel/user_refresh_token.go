package usermodel

import "github.com/go-playground/validator/v10"

type UserRefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (model *UserRefreshToken) Validate() error {
	validate := validator.New()
	return validate.Struct(model)
}
