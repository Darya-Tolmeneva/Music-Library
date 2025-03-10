package main

import (
	"Music_Library/config"
	"Music_Library/internal/database/postgres"
	"Music_Library/internal/router"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func main() {
	cfg := config.Load()

	log := setupLogger(cfg.Env)

	postgres.SetupDatabase(log, cfg)
	r := router.NewRouter()

	if err := r.Run(":8080"); err != nil {
		log.Error("Failed to start server")
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return log
}
