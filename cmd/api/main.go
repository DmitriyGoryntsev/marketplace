package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DmitriyGoryntsev/marketplace/internal/config"
	"github.com/DmitriyGoryntsev/marketplace/internal/transport/http"
	"github.com/DmitriyGoryntsev/marketplace/pkg/logger"
	"github.com/DmitriyGoryntsev/marketplace/pkg/postgres"
	"go.uber.org/zap"
)

func main() {
	ctx, logger, err := logger.New(context.Background())
	signal.NotifyContext(ctx, os.Interrupt, os.Kill, syscall.SIGTERM)
	if err != nil {
		zap.L().Fatal("failed to create logger", zap.Error(err))
	}

	//init config
	cfg, err := config.New()
	if err != nil {
		logger.Fatal(ctx, "failed to read config", zap.Error(err))
	}

	//init db
	db, err := postgres.NewPostgres(cfg.DBConfig)
	if err != nil {
		logger.Fatal(ctx, "failed to connect to database", zap.Error(err))
	}
	defer db.Close(ctx)

	//init repositories
	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)

	//init handlers

	routerConfig := http.NewRouterConfig(cfg)
	r := http.NewRouter(routerConfig, handler)

	go func() {
		logger.Info(ctx, fmt.Sprintf("server is running on port: %s", cfg.HTTPServer.Port))
		r.Run()
	}()

	select {
	case <-ctx.Done():
		logger.Info(ctx, "shutting down the server")
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		r.ShuttingDown()
	}

	logger.Info(ctx, "server shut down")
}
