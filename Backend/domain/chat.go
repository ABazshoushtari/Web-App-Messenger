package domain

type Chat struct {
	BaseModel
	People   []uint64  `json:"people"`
	Messages []Message `gorm:"-" json:"messages,omitempty"`
}

type Message struct {
	BaseModel
	ChatID   uint64 `json:"chat_id"`
	Sender   uint64 `json:"sender"`
	Receiver uint64 `json:"receiver"`
	Content  string `json:"content" validate:"required,max=300"`
}
