package app

import (
	"context"
	"errors"
	"experiment/internal/config"
	"experiment/internal/server"
	router "experiment/internal/ui/api"
	"experiment/pkg/database"
	"experiment/pkg/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	cfg, err := config.Init()
	if err != nil {
		logger.Error(err)
	}

	logger.Info("Config was initialized successfully")

	if cfg.DB.Connection == "mysql" {
		dbConn, err := database.NewMySQLConnection(&cfg.DB)
		if err != nil {
			fmt.Println("Error creating MySQL connection:", err)
			return
		}
		defer dbConn.DB.Close()
	}

	logger.Info("Database was initialized successfully")

	redisConn, err := database.NewRedisConnection(cfg.Redis)
	if err != nil {
		fmt.Println("Error creating and connecting to Redis:", err)
		return
	}
	defer redisConn.Client.Close()

	logger.Info("Redis was initialized successfully")

	redisCacheConn, err := database.NewRedisConnection(cfg.RedisCache)
	if err != nil {
		fmt.Println("Error creating and connecting to Redis Cache:", err)
		return
	}
	defer redisCacheConn.Client.Close()

	logger.Info("Redis Cache was initialized successfully")

	router := router.NewHandler()

	srv := server.NewServer(cfg, router.Init(&cfg.App))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

}
