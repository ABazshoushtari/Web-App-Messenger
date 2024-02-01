package middleware

import (
	"context"
	"github.com/ABazshoushtari/Web-App-Messenger/domain/payloads"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/logger"
	"github.com/ABazshoushtari/Web-App-Messenger/repository"
	"github.com/ABazshoushtari/Web-App-Messenger/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

func NewMiddlewares(svcs *service.Services, repos *repository.Repositories) *Middlewares {
	return &Middlewares{
		svcs:  svcs,
		repos: repos,
		log:   logger.Logger(),
	}
}

type Middlewares struct {
	svcs  *service.Services
	repos *repository.Repositories
	log   *zap.SugaredLogger
}

type CustomContext struct {
	context.Context
	request func() *http.Request
	user    *payloads.UserDTO
}

func (c *CustomContext) GetUser() *payloads.UserDTO {
	return c.user
}

func (m *Middlewares) CustomContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			ctx = &CustomContext{
				Context: ctx,
				request: c.Request,
			}
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
