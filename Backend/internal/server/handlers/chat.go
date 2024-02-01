package handlers

import (
	"github.com/ABazshoushtari/Web-App-Messenger/domain/payloads"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) AddChat() echo.HandlerFunc {
	type request struct {
		payloads.AddChatRequest
	}
	type response struct {
		payloads.AddChatResponse
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return err
		}

		res, err := h.svcs.Chat.AddChat(c.Request().Context(), req.AddChatRequest)
		if err != nil {
			return err
		}
		return c.JSON(200, &response{
			AddChatResponse: *res,
		})
	}
}

func (h *Handlers) IndexChats() echo.HandlerFunc {
	type response struct {
		payloads.IndexChatsResponse
	}
	return func(c echo.Context) error {
		res, err := h.svcs.Chat.IndexChats(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(200, &response{
			IndexChatsResponse: *res,
		})
	}
}
func (h *Handlers) ShowChat() echo.HandlerFunc {
	type response struct {
		payloads.ShowChatResponse
	}
	return func(c echo.Context) error {
		chatID := c.Param("id")
		res, err := h.svcs.Chat.ShowChat(c.Request().Context(), chatID)
		if err != nil {
			return err
		}
		return c.JSON(200, &response{
			ShowChatResponse: *res,
		})

	}
}

func (h *Handlers) DeleteChat() echo.HandlerFunc {
	type response struct {
		payloads.GenericsSuccessFlagResponse
	}
	return func(c echo.Context) error {
		chatID := c.Param("id")
		res, err := h.svcs.Chat.DeleteChat(c.Request().Context(), chatID)
		if err != nil {
			return err
		}
		return c.JSON(200, &response{
			GenericsSuccessFlagResponse: *res,
		})
	}
}
func (h *Handlers) DeleteMessage() echo.HandlerFunc {
	type response struct {
		payloads.GenericsSuccessFlagResponse
	}
	return func(c echo.Context) error {
		chatID := c.Param("chat_id")
		messageID := c.Param("message_id")
		res, err := h.svcs.Chat.DeleteMessage(c.Request().Context(), chatID, messageID)
		if err != nil {
			return err
		}
		return c.JSON(200, &response{
			GenericsSuccessFlagResponse: *res,
		})
	}
}