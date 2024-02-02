package handlers

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) WebsocketHandler(c echo.Context) error {
	upgrader := &websocket.Upgrader{}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error("Error reading message:", err)
			break
		}
		c.Logger().Infof("Received message: %s", string(msg))

		if err := ws.WriteMessage(mt, msg); err != nil {
			c.Logger().Error("Error writing message:", err)
			break
		}
	}

	return nil
}
