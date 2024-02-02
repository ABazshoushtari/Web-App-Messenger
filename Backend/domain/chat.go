package domain

import (
	"github.com/lib/pq"
)

type Chat struct {
	BaseModel
	People   pq.Int64Array `json:"people" gorm:"type:int[]"`
	Messages []Message     `gorm:"-" json:"messages,omitempty"`
}

func (c Chat) TableName() string {
	return "chats"
}

type Message struct {
	BaseModel
	ChatID   uint64 `json:"chat_id"`
	Sender   uint64 `json:"sender"`
	Receiver uint64 `json:"receiver"`
	Content  string `json:"content" validate:"required,max=300"`
}

func (c *Chat) IsParticipant(userID uint64) bool {
	for _, id := range c.People {
		if id == int64(userID) {
			return true
		}
	}
	return false
}
