package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nordew/Strive/internal/config"
	v1 "github.com/nordew/Strive/internal/controller/http/v1"
	"github.com/nordew/Strive/internal/service"
	"github.com/nordew/Strive/internal/storage"
	"github.com/nordew/Strive/pkg/auth"
	"github.com/nordew/Strive/pkg/db/psql"
	"github.com/nordew/Strive/pkg/logger"
	"log"
	"net/http"
)

func MustRun() {
	cfg := config.GetConfig()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pgPool, err := psql.Connect(ctx, cfg.PostgresURL)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer pgPool.Close()

	logger := logger.New()

	userStorage := storage.NewUserStorage(pgPool)
	jwtAuth := auth.NewAuth(logger)
	userService := service.NewUserService(userStorage, jwtAuth, logger)
	router := v1.NewController(userService)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler: router.Init(),
	}

	go func() {
		log.Printf("HTTP server running on port %s", cfg.HTTPPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP server forced to shutdown: %v", err)
	}

	log.Println("HTTP server stopped, cleaning up resources.")
}
