package service

import (
	"context"
	"github.com/ABazshoushtari/Web-App-Messenger/domain/payloads"
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
	AddChat(ctx context.Context, payload payloads.AddChatRequest) (*payloads.AddChatResponse, error)
	IndexChats(ctx context.Context) (*payloads.IndexChatsResponse, error)
	ShowChat(ctx context.Context, chatID uint64) (*payloads.ShowChatResponse, error)
	DeleteChat(ctx context.Context, chatID uint64) (*payloads.GenericsSuccessFlagResponse, error)
	DeleteMessage(ctx context.Context, chatID uint64, messageID uint64) (*payloads.GenericsSuccessFlagResponse, error)
}
type User interface {
	AuthRegister(ctx context.Context, payload payloads.UserRegisterRequest) (*payloads.UserRegisterResponse, error)
	AuthLogin(ctx context.Context, payload payloads.UserLoginRequest) (*payloads.UserLoginResponse, error)
	ShowUser(ctx context.Context, userID uint64) (*payloads.UserShowResponse, error)
	IndexUser(ctx context.Context, keyword string) (*payloads.UserShowResponse, error)
	UpdateUser(ctx context.Context, userID uint64, payload payloads.UserUpdateRequest) (*payloads.GenericsSuccessFlagResponse, error)
	DeleteUser(ctx context.Context, userID uint64) (*payloads.GenericsSuccessFlagResponse, error)
}
type Contact interface {
	ShowContacts(ctx context.Context, userID uint64) (*payloads.ShowContactsResponse, error)
	AddContact(ctx context.Context, userID uint64, payload payloads.AddContactRequest) (*payloads.AddContactResponse, error)
	DeleteContact(ctx context.Context, userID uint64, contactID uint64) (*payloads.GenericsSuccessFlagResponse, error)
}
