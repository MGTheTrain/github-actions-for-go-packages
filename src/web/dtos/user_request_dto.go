package dtos

import (
	"github.com/go-playground/validator/v10"
)

type UserRequestDto struct {
	UserName     *string `validate:"required,max=50"`
	UserPassword *string `validate:"required,min=10"`
	Email        *string `validate:"required"`
}

func (u *UserRequestDto) Validate() []string {
	validate := validator.New()
	var validationErrors []string

	err := validate.Struct(u)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, e.Field()+" "+e.Tag())
		}
	}
	return validationErrors
}
