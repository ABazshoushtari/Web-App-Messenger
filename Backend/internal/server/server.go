package server

import (
	"github.com/ABazshoushtari/Web-App-Messenger/internal/config"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/infra"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/server/handlers"
	"github.com/ABazshoushtari/Web-App-Messenger/repository"
	"github.com/ABazshoushtari/Web-App-Messenger/service"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"os"
)

func Start() error {
	e := echo.New()

	e.Use(echoMiddleware.Logger())
	e.HTTPErrorHandler = ErrorHandler()

	db, err := infra.NewGORMConnection(config.GetConfig().DB.GetDSN())
	if err != nil {
		return err
	}
	repos := repository.NewRepositories(db)

	svcs := service.NewServices(repos)
	handler := handlers.New(svcs)
	routes(e, handler)
	return e.Start(":" + os.Getenv("APP_PORT"))
}
