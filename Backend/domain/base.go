package domain

import (
	"context"
	"github.com/ABazshoushtari/Web-App-Messenger/domain/payloads"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/server/middleware"
	"time"
)

type BaseModel struct {
	ID        uint64    `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func GetUserDTO(ctx context.Context) *payloads.UserDTO {
	if ctx == nil {
		return nil
	}
	if c, ok := ctx.(middleware.CustomContext); ok {
		return c.GetUser()
	}
	return nil
}
