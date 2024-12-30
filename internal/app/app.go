package app

import (
	"context"
	"github.com/nordew/Strive/internal/config"
	"github.com/nordew/Strive/internal/service"
	"github.com/nordew/Strive/internal/storage"
	"github.com/nordew/Strive/pkg/db/psql"
	"github.com/nordew/Strive/pkg/logger"
	"log"
)

func MustRun() {
	cfg := config.GetConfig()

	pgConn, err := psql.Connect(context.Background(), cfg.PostgresURL)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	logger := logger.New()

	userStorage := storage.NewUserStorage(pgConn)

	userService := service.NewUserService(userStorage, logger)
}
