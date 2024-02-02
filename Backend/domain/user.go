package domain

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	BaseModel
	FirstName   string `json:"first_name" validate:"required,min=1"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number" validate:"required,number,min=11,max=12"`
	Username    string `json:"username" validate:"required,min=5,alphanum"`
	Password    string `json:"password"`
	Image       string `json:"image"`
	Bio         string `json:"bio" validate:"max=100"`
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

	if err != nil {
		return err
	}
	return nil
}

func (u User) ToDTO() *UserDTO {
	return &UserDTO{
		User:     u,
		Password: "********",
	}
}
