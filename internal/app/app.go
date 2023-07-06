package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/begenov/api-gateway/internal/client"
	"github.com/begenov/api-gateway/internal/config"
	httpHandler "github.com/begenov/api-gateway/internal/delivery/http"
	"github.com/begenov/api-gateway/internal/repository"
	"github.com/begenov/api-gateway/internal/server"
	"github.com/begenov/api-gateway/internal/service"
	"github.com/begenov/api-gateway/pkg/database"
	"github.com/begenov/api-gateway/pkg/logger"
)

func Run(cfg *config.Config) error {
	logger := logger.CreateLogger(cfg.Log)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db, err := database.Open(ctx, cfg.POSTGRES)
	if err != nil {
		return err
	}

	registergRPCClient, err := client.NewRegisterServiceClient(cfg.RegisterServiceAddr)
	if err != nil {
		return err
	}

	repository := repository.NewRepository(db)

	service := service.NewService(repository, logger)

	httpHandler := httpHandler.NewHandler(service, logger, registergRPCClient)

	srv := server.NewServer(&cfg.Log.HTTP, httpHandler.Init())

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Sugar().Infof("Server started: %s", cfg.Log.HTTP.Port)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("failed to stop server: %v", err)
	}

	return nil
}
