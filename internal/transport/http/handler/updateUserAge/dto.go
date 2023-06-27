package updateHandler

import (
	validation "github.com/go-ozzo/ozzo-validation"
	govalidator "github.com/go-ozzo/ozzo-validation/is"
)

type DTO struct {
	Id  string `form:"id" json:"id"`
	Age string `json:"new_age"`
}

func (dto DTO) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Id, validation.Required),

		validation.Field(&dto.Age, validation.Required),
		validation.Field(&dto.Age, govalidator.Digit),
	)
}
