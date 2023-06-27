package createHandler

import (
	validation "github.com/go-ozzo/ozzo-validation"
	govalidator "github.com/go-ozzo/ozzo-validation/is"
)

type DTO struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func (d DTO) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name, validation.Required),
		validation.Field(&d.Name, govalidator.Alpha),

		validation.Field(&d.Age, validation.Required),
		validation.Field(&d.Age, govalidator.Digit),
	)
}
