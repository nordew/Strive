package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nordew/Strive/internal/bots"
	"github.com/nordew/Strive/internal/config"
	"github.com/nordew/Strive/internal/controller/http/v1"
	"github.com/nordew/Strive/internal/service"
	"github.com/nordew/Strive/internal/storage"
	"github.com/nordew/Strive/pkg/auth"
	"github.com/nordew/Strive/pkg/db/psql"
	"github.com/nordew/Strive/pkg/logger"
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

	// Start HTTP server in a separate goroutine
	go func() {
		log.Printf("HTTP server running on port %d", cfg.HTTPPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// Start the Telegram bot in a separate goroutine
	go func() {
		log.Println("Initializing bots bot...")
		if err := bots.InitBot(cfg.BOTToken, cfg.WebAppURL); err != nil {
			log.Fatalf("failed to init bots bot: %v", err)
		}
	}()

	// Await shutdown signal
	<-ctx.Done()

	log.Println("Shutting down gracefully...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP server forced to shutdown: %v", err)
	}

	log.Println("HTTP server stopped, cleaning up resources.")
}
