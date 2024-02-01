package repository

import (
	"errors"
	"github.com/ABazshoushtari/Web-App-Messenger/domain"
	"gorm.io/gorm"
	"strconv"
)

type chatRepository struct {
	db *gorm.DB
}

func newChatRepository(db *gorm.DB) *chatRepository {
	return &chatRepository{
		db: db,
	}
}

func (c *chatRepository) Create(chat *domain.Chat) error {
	return c.db.Create(&chat).Error
}

func (c *chatRepository) GetByID(chatID uint64, chat *domain.Chat) error {
	if err := c.db.Where("id = ?", chatID).First(&chat).Error; err != nil {
		return err
	}
	messages, err := c.getMessages(chatID)
	chat.Messages = messages
	if errors.Is(err, gorm.ErrRecordNotFound) || err == nil {
		return nil
	}
	return err
}

func (c *chatRepository) GetByUserID(userID uint64) ([]domain.Chat, error) {
	chats := []domain.Chat{}
	err := c.db.Where(strconv.Itoa(int(userID)) + "= ANY(people)").Find(&chats).Error
	return chats, err
} //TODO: check if this function actually works

func (c *chatRepository) Delete(chatID uint64) error {
	return c.db.Where("id = ?", chatID).Delete(&domain.Chat{}).Error
}

func (c *chatRepository) DeleteMessage(chatID uint64, messageID uint64) error {
	return c.db.Where("chat_id = ? AND id = ?", chatID, messageID).Delete(&domain.Message{}).Error
}
func (c *chatRepository) DeleteAllMessages(chatID uint64) error {
	return c.db.Where("chat_id = ?", chatID).Delete(&domain.Message{}).Error
}
func (c *chatRepository) getMessages(chatID uint64) ([]domain.Message, error) {
	messages := []domain.Message{}
	err := c.db.Where("chat_id = ?", chatID).Find(&messages).Error
	return messages, err
}

func (c *chatRepository) GetByParticipants(firstUser uint64, secondUser uint64) error {
	return c.db.Where("ARRAY[?,?] <@ people", firstUser, secondUser).Find(&domain.Chat{}).Error
}

func (c *chatRepository) DeleteUserFromChats(userID uint64) error {
	return c.db.Exec("SELECT ARRAY_REMOVE(people, ?) FROM chats", userID).Error
}
