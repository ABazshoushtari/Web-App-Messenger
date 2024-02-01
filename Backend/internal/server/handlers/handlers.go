package handlers

import (
	"github.com/ABazshoushtari/Web-App-Messenger/internal/logger"
	"github.com/ABazshoushtari/Web-App-Messenger/service"
	"go.uber.org/zap"
)

type Handlers struct {
	svcs *service.Services
	log  *zap.SugaredLogger
}

func New(svcs *service.Services) *Handlers {
	return &Handlers{
		svcs: svcs,
		log:  logger.Logger(),
	}
}
