package handlers

import (
	"errors"
	"github.com/ABazshoushtari/Web-App-Messenger/domain/payloads"
	"github.com/labstack/echo/v4"
	"regexp"
	"strconv"
)

func (h *Handlers) ShowUser() echo.HandlerFunc {
	type response struct {
		payloads.UserShowResponse
	}
	return func(c echo.Context) error {
		userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return errors.New("invalid user id")
		}
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
		payloads.UserShowResponse
	}
	return func(c echo.Context) error {
		key := c.QueryParam("keyword")
		if isMatch, _ := regexp.MatchString("^[a-zA-Z0-9]*$", key); !isMatch {
			return errors.New("invalid keyword. only numbers and letters are allowed")
		}
		res, err := h.svcs.User.IndexUser(c.Request().Context(), key)
		if res.UserDTO == nil {
			return errors.New("not Found")
		}
		if err != nil {
			return err
		}
		return c.JSON(200, response{
			UserShowResponse: *res,
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

		userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return errors.New("invalid user id")
		}
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
		userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return errors.New("invalid user id")
		}
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
		userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return errors.New("invalid user id")
		}
		res, err := h.svcs.Contact.ShowContacts(c.Request().Context(), userID)
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
		userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return errors.New("invalid user id")
		}
		res, err := h.svcs.Contact.AddContact(c.Request().Context(), userID, req.AddContactRequest)
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
		userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return errors.New("invalid user id")
		}
		contactID, err := strconv.ParseUint(c.Param("contact_id"), 10, 64)
		if err != nil {
			return errors.New("invalid contact id")
		}
		res, err := h.svcs.Contact.DeleteContact(c.Request().Context(), userID, contactID)
		if err != nil {
			return err
		}
		return c.JSON(200, response{
			GenericsSuccessFlagResponse: *res,
		})
	}
}
