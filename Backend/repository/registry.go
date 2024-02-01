package repository

import (
	"github.com/ABazshoushtari/Web-App-Messenger/domain"
	"gorm.io/gorm"
	"mime/multipart"
)

type Repositories struct {
	User    User
	Chat    Chat
	Contact Contact
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:    newUserRepository(db),
		Chat:    newChatRepository(db),
		Contact: newContactRepository(db),
	}
}

type User interface {
	Register(user *domain.User) error
	GetByUsername(username string, user *domain.User) error
	GetByPhoneNumber(phoneNumber string, user *domain.User) error
	GetByID(id uint64, user *domain.User) error
	Update(user *domain.User) error
	Delete(user *domain.User) error
	GetByKey(value string, user *domain.User) error
	CheckExisting(username string, phoneNumber string) error
	SetImage(user *domain.User, image *multipart.FileHeader) error
}

type Chat interface {
	Create(chat *domain.Chat) error
	GetByID(chatID uint64, chat *domain.Chat) error
	GetByUserID(userID uint64) ([]domain.Chat, error)
	GetByParticipants(firstUser uint64, secondUser uint64) error
	Delete(chatID uint64) error
	DeleteAllMessages(chatID uint64) error
	DeleteMessage(chatID uint64, messageID uint64) error
}

type Contact interface {
	Upsert(contact *domain.Contact) error
	GetByUserID(userID uint64) ([]domain.Contact, error)
	Delete(userID uint64, contactID uint64) error
}
