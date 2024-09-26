package gingonic

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/shakew3ll/Telegram-Bot-Service.git/pkg/logging"
)

type Server struct {
	logger *logging.Logger
	router *gin.Engine
	bindIp string
	port   int
}

func NewApp(
	logger *logging.Logger,
	router *gin.Engine,
	bindIp string,
	port int,
) *Server {
	return &Server{
		logger: logger,
		router: router,
		bindIp: bindIp,
		port:   port,
	}
}

func (a *Server) MustRun() {
	defer func() {
		if r := recover(); r != nil {
			a.logger.Fatalf("Application panicked: %v", r)
		}
	}()

	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *Server) Run() error {
	a.logger.Info("Starting HTTP server...")

	addr := fmt.Sprintf("%s:%d", a.bindIp, a.port)
	a.logger.Infof("HTTP server is listening on host: %s", addr)

	if err := a.router.Run(addr); err != nil {
		a.logger.Fatalf("Failed to start HTTP server due to an error: %v", err)
		return err
	}

	return nil
}

func (a *Server) Stop(ctx context.Context) {
	done := make(chan struct{})

	go func() {
		defer close(done)
		a.logger.Info("Stopping HTTP server...")
	}()

	select {
	case <-done:
		a.logger.Info("Server gracefully stopped.")
	case <-ctx.Done():
		a.logger.Warn("Server shutdown timed out, forcing exit.")
	}
}
