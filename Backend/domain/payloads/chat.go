package payloads

import "github.com/ABazshoushtari/Web-App-Messenger/domain"

type AddChatRequest struct {
	ParticipantID uint64 `json:"participant_id"`
}

type AddChatResponse struct {
	Chat domain.Chat `json:"chat"`
}

type IndexChatsResponse struct {
	Chats []domain.Chat `json:"chats"`
}

type ShowChatResponse struct {
	Chat domain.Chat `json:"chat"`
}
