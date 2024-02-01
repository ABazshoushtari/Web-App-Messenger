package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if err != nil {
			if httpErr, ok := err.(*echo.HTTPError); ok {
				c.JSON(httpErr.Code, httpErr.Message)
				return
			}
			c.JSON(http.StatusInternalServerError, map[string]string{
				"Code": "unexpected_error",
				"Msg":  err.Error(),
			})
			return
		}
	}
}
