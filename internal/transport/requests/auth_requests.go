package requests

import "github.com/go-playground/validator/v10"

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (req UserLogin) Validate() error {

	return validator.New().Struct(
		&req,
	)
}
