package domain

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type User struct {
	BaseModel
	FirstName   string `json:"first_name" validate:"required,min=1"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number" validate:"required,number"`
	Username    string `json:"username" validate:"required,min=5,alphanum"`
	Password    string `json:"password"`
	Image       string `json:"image"`
	Bio         string `json:"bio" validate:"required,max=100"`
}
type UserDTO struct {
	User
	Password string `default:"" json:"password"`
}

var userValidator *validator.Validate

func init() {
	userValidator = validator.New()
}

func (u User) Validate() error {
	err := userValidator.Struct(u)
	validationErr := validator.ValidationErrors{}

	if err != nil && errors.Is(err, &validationErr) {
		return validationErr
	}
	return nil
}

func (u User) ToDTO() *UserDTO {
	return &UserDTO{
		User:     u,
		Password: "********",
	}
}
