package deleteHandler

import (
	validation "github.com/go-ozzo/ozzo-validation"
	govalidator "github.com/go-ozzo/ozzo-validation/is"
)

type DTO struct {
	Id string `json:"target_id"`
}

func (dto DTO) Validate() error {

	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Id, validation.Required),
		validation.Field(&dto.Id, govalidator.Digit),
	)

}
