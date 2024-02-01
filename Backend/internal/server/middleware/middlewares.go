package middleware

import (
	"github.com/ABazshoushtari/Web-App-Messenger/internal/logger"
	"github.com/ABazshoushtari/Web-App-Messenger/repository"
	"github.com/ABazshoushtari/Web-App-Messenger/service"
	"go.uber.org/zap"
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
