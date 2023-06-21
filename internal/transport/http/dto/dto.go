package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	govalidator "github.com/go-ozzo/ozzo-validation/is"
)

// var (
// 	ErrValidateWrongName = errors.New("name: must contain English letters only.")
// 	ErrValidateWrongAge  = errors.New("age: must contain digits only.")
// )

type MakeFriendDTO struct {
	Target string `json:"target_id"`
	Source string `json:"source_id"`
}

func (dto MakeFriendDTO) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Target, validation.Required),
		validation.Field(&dto.Target, govalidator.Digit),

		validation.Field(&dto.Source, validation.Required),
		validation.Field(&dto.Source, govalidator.Digit),
	)

}

type DeleteUserDto struct {
	Id string `json:"target_id"`
}

func (dto DeleteUserDto) Validate() error {

	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Id, validation.Required),
		validation.Field(&dto.Id, govalidator.Digit),
	)

}

type UpdateUserAgeDTO struct {
	Id  uint64 `form:"id"`
	Age string `json:"new age"`
}

func (dto UpdateUserAgeDTO) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Id, validation.Required),

		validation.Field(&dto.Age, validation.Required),
		validation.Field(&dto.Age, govalidator.Digit),
	)
}

type CreateUserDTO struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func (dto CreateUserDTO) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Name, validation.Required),
		validation.Field(&dto.Name, govalidator.Alpha),

		validation.Field(&dto.Age, validation.Required),
		validation.Field(&dto.Age, govalidator.Digit),
	)

}

type GetFriendsDto struct {
	Id uint64 `json:"id"`
}

func (dto GetFriendsDto) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Id, validation.Required),
	)

}
