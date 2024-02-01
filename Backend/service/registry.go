package service

import (
	"github.com/ABazshoushtari/Web-App-Messenger/repository"
	"github.com/ABazshoushtari/Web-App-Messenger/service/chat"
	"github.com/ABazshoushtari/Web-App-Messenger/service/contact"
	"github.com/ABazshoushtari/Web-App-Messenger/service/user"
)

type Services struct {
	Chat    Chat
	User    User
	Contact Contact
}

func NewServices(repo *repository.Repositories) *Services {
	return &Services{
		Chat:    chat.New(repo),
		User:    user.New(repo),
		Contact: contact.New(repo),
	}
}

type Chat interface {
}
type User interface {
}
type Contact interface {
}
