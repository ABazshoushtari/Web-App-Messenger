package repository

import (
	"github.com/ABazshoushtari/Web-App-Messenger/domain"
	"gorm.io/gorm"
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
	GetByKey(key string, value string, user *domain.User) error
}

type Chat interface {
	Create(chat *domain.Chat) error
	GetByID(chatID uint64, chat *domain.Chat) error
	GetByUserID(userID uint64) ([]domain.Chat, error)
}

type Contact interface {
}
