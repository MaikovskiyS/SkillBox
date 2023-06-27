package createFriendship

import (
	validation "github.com/go-ozzo/ozzo-validation"
	govalidator "github.com/go-ozzo/ozzo-validation/is"
)

type DTO struct {
	Target string `json:"target_id"`
	Source string `json:"source_id"`
}

func (dto DTO) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Target, validation.Required),
		validation.Field(&dto.Target, govalidator.Digit),

		validation.Field(&dto.Source, validation.Required),
		validation.Field(&dto.Source, govalidator.Digit),
	)

}
