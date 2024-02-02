package server

import (
	"github.com/ABazshoushtari/Web-App-Messenger/internal/config"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/infra"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/server/handlers"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/server/middleware"
	"github.com/ABazshoushtari/Web-App-Messenger/repository"
	"github.com/ABazshoushtari/Web-App-Messenger/service"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Start() error {
	e := echo.New()

	db, err := infra.NewGORMConnection(config.GetConfig().DB.GetDSN())
	if err != nil {
		return err
	}
	repos := repository.NewRepositories(db)
	svcs := service.NewServices(repos)
	middlewares := middleware.NewMiddlewares(svcs, repos)
	handler := handlers.New(svcs)
	e.Use(middlewares.CustomContext())
	e.Use(echoMiddleware.Logger())
	e.HTTPErrorHandler = ErrorHandler()

	routes(e, handler, middlewares)
	return e.Start(":" + config.GetConfig().AppPort)
}
