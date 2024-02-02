package payloads

import (
	"github.com/ABazshoushtari/Web-App-Messenger/domain"
	"mime/multipart"
)

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
	User domain.UserDTO `json:"user"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserUpdateRequest struct {
	Username  string                `form:"username"`
	Password  string                `form:"password"`
	FirstName string                `form:"first_name"`
	LastName  string                `form:"last_name"`
	Phone     string                `form:"phone"`
	Image     *multipart.FileHeader `form:"image"`
	Bio       string                `form:"bio"`
}

type UserShowResponse struct {
	*domain.UserDTO `json:"user"`
}
type UserIndexResponse struct {
	Users []domain.UserDTO `json:"users"`
}
