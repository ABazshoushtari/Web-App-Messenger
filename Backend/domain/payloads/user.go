package payloads

import (
	"github.com/ABazshoushtari/Web-App-Messenger/domain"
	"mime/multipart"
)

type UserDTO struct {
	domain.User
	Password string `default:"" json:"password"`
}
type UserRegisterRequest struct {
	Username  string                `form:"username"`
	Password  string                `form:"password"`
	FirstName string                `form:"first_name"`
	LastName  string                `form:"last_name,omitempty"`
	Phone     string                `form:"phone"`
	Image     *multipart.FileHeader `form:"image,omitempty"`
	Bio       string                `form:"bio,omitempty"`
}

type UserRegisterResponse struct {
	User UserDTO `json:"user"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserUpdateRequest struct {
	Username  string                `form:"username,omitempty"`
	Password  string                `form:"password,omitempty"`
	FirstName string                `form:"first_name,omitempty"`
	LastName  string                `form:"last_name,omitempty"`
	Phone     string                `form:"phone,omitempty"`
	Image     *multipart.FileHeader `form:"image,omitempty"`
	Bio       string                `form:"bio,omitempty"`
}

type UserShowResponse struct {
	*UserDTO `json:"user"`
}
type UserIndexResponse struct {
	Users []UserDTO `json:"users"`
}
