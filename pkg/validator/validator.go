package validator

import "github.com/go-playground/validator/v10"

type UserLogin struct {
	Username string `json:"username" validate:"required,min=3,max=30"`
	Password string `json:"password" validate:"required,min=6"`
}

func ValidateLogin(user UserLogin) error {
	validate := validator.New()
	return validate.Struct(user)
}
