package handlers

import (
	"errors"
	"github.com/ABazshoushtari/Web-App-Messenger/domain/payloads"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) ShowUser() echo.HandlerFunc {
	type response struct {
		payloads.UserShowResponse
	}
	return func(c echo.Context) error {
		userID := c.Param("id")
		res, err := h.svcs.User.ShowUser(c.Request().Context(), userID)
		if err != nil {
			return err
		}
		return c.JSON(200, response{
			UserShowResponse: *res,
		})
	}
}

func (h *Handlers) IndexUser() echo.HandlerFunc {
	type response struct {
		payloads.UserIndexResponse
	}
	return func(c echo.Context) error {
		key := c.QueryParam("keyword")
		res, err := h.svcs.User.IndexUser(c.Request().Context(), key)
		if len(res.Users) == 0 {
			return errors.New("not Found")
		}
		if err != nil {
			return err
		}
		return c.JSON(200, response{
			UserIndexResponse: *res,
		})
	}
}

func (h *Handlers) UpdateUser() echo.HandlerFunc {
	type request struct {
		payloads.UserUpdateRequest
	}
	type response struct {
		payloads.GenericsSuccessFlagResponse
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return errors.New("invalid form request")
		}
		//img, err := c.FormFile("image")
		//if err != nil {
		//	return errors.New("invalid user image")
		//}
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

		userID := c.Param("id")
		res, err := h.svcs.User.UpdateUser(c.Request().Context(), userID, req.UserUpdateRequest)
		if err != nil {
			return err
		}
		return c.JSON(200, response{
			GenericsSuccessFlagResponse: *res,
		})
	}
}

func (h *Handlers) DeleteUser() echo.HandlerFunc {
	type response struct {
		payloads.GenericsSuccessFlagResponse
	}
	return func(c echo.Context) error {
		userID := c.Param("id")
		res, err := h.svcs.User.DeleteUser(c.Request().Context(), userID)
		if err != nil {
			return err
		}
		return c.JSON(200, response{
			GenericsSuccessFlagResponse: *res,
		})
	}
}

func (h *Handlers) ShowContacts() echo.HandlerFunc {
	type response struct {
		payloads.ShowContactsResponse
	}
	return func(c echo.Context) error {
		userID := c.Param("id")
		res, err := h.svcs.User.ShowContacts(c.Request().Context(), userID)
		if err != nil {
			return err
		}
		return c.JSON(200, response{
			ShowContactsResponse: *res,
		})
	}
}

func (h *Handlers) AddContact() echo.HandlerFunc {
	type request struct {
		payloads.AddContactRequest
	}
	type response struct {
		payloads.AddContactResponse
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return errors.New("invalid request")
		}
		userID := c.Param("id")
		res, err := h.svcs.User.AddContact(c.Request().Context(), userID, req.AddContactRequest)
		if err != nil {
			return err
		}
		return c.JSON(200, response{
			AddContactResponse: *res,
		})
	}
}

func (h *Handlers) DeleteContact() echo.HandlerFunc {
	type response struct {
		payloads.GenericsSuccessFlagResponse
	}
	return func(c echo.Context) error {
		userID := c.Param("id")
		contactID := c.Param("contact_id")
		res, err := h.svcs.User.DeleteContact(c.Request().Context(), userID, contactID)
		if err != nil {
			return err
		}
		return c.JSON(200, response{
			GenericsSuccessFlagResponse: *res,
		})
	}
}
