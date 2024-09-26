package gingonic

import (
	"github.com/gin-gonic/gin"

	"github.com/shakew3ll/Telegram-Bot-Service.git/pkg/logging"
)

type Application struct {
	Server *Server
}

func New(
	logger *logging.Logger,
	bindIp string,
	port int,
) *Application {
	logger.Info("Initializing router...")
	router := gin.Default()

	ginApp := NewApp(logger, router, bindIp, port)

	return &Application{
		Server: ginApp,
	}
}
