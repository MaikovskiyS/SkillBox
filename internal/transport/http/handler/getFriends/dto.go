package getFriendsHandler

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type DTO struct {
	Id string `form:"id"`
}

func (dto DTO) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Id, validation.Required),
	)

}
