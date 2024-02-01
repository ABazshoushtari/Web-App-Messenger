package server

import (
	"github.com/ABazshoushtari/Web-App-Messenger/internal/server/handlers"
	"github.com/labstack/echo/v4"
)

func routes(e *echo.Echo, h *handlers.Handlers) {
	api := e.Group("/api")
	api.POST("/register", h.AuthRegister())
	api.POST("/login", h.AuthLogin())

	user := api.Group("/users")
	user.GET("/:id", h.ShowUser())
	user.PATCH("/:id", h.UpdateUser())
	user.DELETE("/:id", h.DeleteUser())
	user.GET("", h.IndexUser()) // query by ey

	user.GET("/:id/contacts", h.ShowContacts())
	user.POST("/:id/contacts", h.AddContact())
	user.DELETE("/:id/contacts/:contact_id", h.DeleteContact())

	chat := api.Group("/chats")
	chat.POST("", h.AddChat())
	chat.GET("", h.IndexChats())
	chat.GET("/:id", h.ShowChat())
	chat.DELETE("/:id", h.DeleteChat())
	chat.DELETE("/:id/messages/:message_id", h.DeleteMessage())

}