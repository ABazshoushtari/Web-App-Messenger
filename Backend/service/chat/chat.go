package chat

import (
	"context"
	"errors"
	"github.com/ABazshoushtari/Web-App-Messenger/domain"
	"github.com/ABazshoushtari/Web-App-Messenger/domain/payloads"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/logger"
	"github.com/ABazshoushtari/Web-App-Messenger/repository"
	"gorm.io/gorm"
)

type Chat struct {
	repos *repository.Repositories
}

func New(repos *repository.Repositories) *Chat {
	return &Chat{
		repos: repos,
	}
}

func (c *Chat) AddChat(ctx context.Context, payload payloads.AddChatRequest) (*payloads.AddChatResponse, error) {
	user := domain.GetUserDTO(ctx)
	if err := c.repos.Chat.GetByParticipants(user.ID, payload.ParticipantID); !errors.Is(err, gorm.ErrRecordNotFound) {
		if err == nil {
			return nil, errors.New("chat already exists")
		}
		logger.Logger().Errorw("error while checking if chat exists", "error", err)
		return nil, errors.New("error while checking if chat exists")
	}
	chat := domain.Chat{
		People: []uint64{user.ID, payload.ParticipantID},
	}
	if err := c.repos.Chat.Create(&chat); err != nil {
		logger.Logger().Errorw("error while creating chat", "error", err)
		return nil, errors.New("error while creating chat")
	}
	return &payloads.AddChatResponse{
		Chat: chat,
	}, nil
}

func (c *Chat) IndexChats(ctx context.Context) (*payloads.IndexChatsResponse, error) {
	user := domain.GetUserDTO(ctx)
	if user == nil {
		return nil, errors.New("invalid jwt user")
	}
	chats, err := c.repos.Chat.GetByUserID(user.ID)
	if err != nil {
		logger.Logger().Errorw("error while getting chats", "error", err)
		return nil, errors.New("error while getting chats")
	}
	return &payloads.IndexChatsResponse{
		Chats: chats,
	}, nil
}

func (c *Chat) ShowChat(ctx context.Context, chatID uint64) (*payloads.ShowChatResponse, error) {
	user := domain.GetUserDTO(ctx)
	if user == nil {
		return nil, errors.New("invalid jwt user")
	}
	chat := domain.Chat{}
	if err := c.repos.Chat.GetByID(chatID, &chat); err != nil {
		logger.Logger().Errorw("error while getting chat", "error", err)
		return nil, errors.New("error while getting chat")
	}
	if !chat.IsParticipant(user.ID) {
		return nil, errors.New("you are not a participant of this chat")
	}
	return &payloads.ShowChatResponse{
		Chat: chat,
	}, nil
}

func (c *Chat) DeleteChat(ctx context.Context, chatID uint64) (*payloads.GenericsSuccessFlagResponse, error) {
	user := domain.GetUserDTO(ctx)
	if user == nil {
		return nil, errors.New("invalid jwt user")
	}
	chat := domain.Chat{}
	if err := c.repos.Chat.GetByID(chatID, &chat); err != nil {
		logger.Logger().Errorw("error while getting chat", "error", err)
		return nil, errors.New("error while getting chat")
	}
	if !chat.IsParticipant(user.ID) {
		return nil, errors.New("you are not a participant of this chat")
	}
	if err := c.repos.Chat.Delete(chatID); err != nil {
		logger.Logger().Errorw("error while deleting chat", "error", err)
		return nil, errors.New("error while deleting chat")
	}
	_ = c.repos.Chat.DeleteAllMessages(chatID)

	return &payloads.GenericsSuccessFlagResponse{
		Successful: true,
		Message:    "chat deleted successfully",
	}, nil
}

func (c *Chat) DeleteMessage(ctx context.Context, chatID uint64, messageID uint64) (*payloads.GenericsSuccessFlagResponse, error) {
	user := domain.GetUserDTO(ctx)
	if user == nil {
		return nil, errors.New("invalid jwt user")
	}
	chat := domain.Chat{}
	if err := c.repos.Chat.GetByID(chatID, &chat); err != nil {
		logger.Logger().Errorw("error while getting chat", "error", err)
		return nil, errors.New("error while getting chat")
	}
	if !chat.IsParticipant(user.ID) {
		return nil, errors.New("you are not a participant of this chat")
	}
	if err := c.repos.Chat.DeleteMessage(chatID, messageID); err != nil {
		logger.Logger().Errorw("error while deleting message", "error", err)
		return nil, errors.New("error while deleting message")
	}
	return &payloads.GenericsSuccessFlagResponse{
		Successful: true,
		Message:    "message deleted successfully",
	}, nil
}
