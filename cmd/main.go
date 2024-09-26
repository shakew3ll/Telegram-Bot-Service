package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/shakew3ll/Telegram-Bot-Service.git/config"
	"github.com/shakew3ll/Telegram-Bot-Service.git/infrastructure/gingonic"
	"github.com/shakew3ll/Telegram-Bot-Service.git/pkg/logging"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load the config due an error: %v", err.Error())
	}

	logger, err := logging.New(cfg)
	if err != nil {
		log.Fatalf("Failed to configure logger due an error: %v", err.Error())
	}
	logger.Info("Logger connected successfully.")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	app := gingonic.New(logger, cfg.Listen.Host, cfg.Listen.Port)
	go app.Server.MustRun()

	<-ctx.Done()
	logger.Info("Received shutdown signal, shutting down application...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Timeout.Value)
	defer cancel()

	app.Server.Stop(shutdownCtx)
}
