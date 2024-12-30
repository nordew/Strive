package app

import (
	"context"
	"github.com/nordew/Strive/internal/config"
	v1 "github.com/nordew/Strive/internal/controller/http/v1"
	"github.com/nordew/Strive/internal/service"
	"github.com/nordew/Strive/internal/storage"
	"github.com/nordew/Strive/pkg/auth"
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

	jwtAuth := auth.NewAuth(logger)

	userService := service.NewUserService(userStorage, jwtAuth, logger)

	router := v1.NewController(userService)

	router.InitAndRun(cfg.HTTPPort)
}
