package domain

import (
	"context"
	"net/http"
	"time"
)

type BaseModel struct {
	ID        uint64    `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CustomContext struct {
	context.Context
	Request func() *http.Request
	User    *UserDTO
}

func (c *CustomContext) GetUser() *UserDTO {
	return c.User
}
func GetUserDTO(ctx context.Context) *UserDTO {
	if ctx == nil {
		return nil
	}
	if c, ok := ctx.(*CustomContext); ok {
		return c.GetUser()
	}
	return nil
}
