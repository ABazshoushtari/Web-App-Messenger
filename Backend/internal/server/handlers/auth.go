package handlers

import (
	"errors"
	"github.com/ABazshoushtari/Web-App-Messenger/domain/payloads"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handlers) AuthRegister() echo.HandlerFunc {
	type request struct {
		payloads.UserRegisterRequest
	}
	type response struct {
		payloads.UserRegisterResponse
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return errors.New("invalid form request")
		}
		img, err := c.FormFile("image")
		if errors.Is(err, http.ErrMissingFile) || err == nil {
			req.Image = img
			if req.Image != nil {
				if req.Image.Size > 1000000 {
					return errors.New("image size too large")
				}
			}
		} else {
			return errors.New("invalid image")
		}

		//username := c.FormValue("username")
		//password := c.FormValue("password")
		//firstName := c.FormValue("first_name")
		//lastName := c.FormValue("last_name")
		//phone := c.FormValue("phone")
		//bio := c.FormValue("bio")
		//
		//req.Username = username
		//req.Password = password
		//req.FirstName = firstName
		//req.LastName = lastName
		//req.Phone = phone
		//req.Image = img
		//req.Bio = bio

		res, err := h.svcs.User.AuthRegister(c.Request().Context(), req.UserRegisterRequest)
		if err != nil {
			return err
		}
		return c.JSON(200, response{
			UserRegisterResponse: *res,
		})
	}
}

func (h *Handlers) AuthLogin() echo.HandlerFunc {
	type request struct {
		payloads.UserLoginRequest
	}
	type response struct {
		payloads.UserLoginResponse
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return errors.New("invalid request body")
		}

		res, err := h.svcs.User.AuthLogin(c.Request().Context(), req.UserLoginRequest)
		if err != nil {
			return err
		}
		return c.JSON(200, response{
			UserLoginResponse: *res,
		})
	}
}
