package models

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	UserID          *string `validate:"required"`
	UserName        *string `validate:"required,max=50"`
	UserPassword    *string `validate:"required,min=10"`
	Email           *string `validate:"required"`
	DateTimeCreated *string `validate:"required"`
	DateTimeUpdated *string `validate:"required"`
}

func (u *User) Validate() []string {
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
