package postgres

import (
	"Music_Library/config"
	"Music_Library/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

var DB *gorm.DB

func SetupDatabase(log *slog.Logger, cfg *config.Config) {
	host := cfg.Storage.Host
	port := cfg.Storage.Port
	user := cfg.Storage.Username
	name := cfg.Storage.Database
	password := cfg.Storage.Password

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		host, user, name, password, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Failed to connect to database")
	} else {
		log.Info("Database connection established")
	}

	err = DB.AutoMigrate(&models.Song{}, &models.Lyric{})
	if err != nil {
		log.Error("Failed to migrate database")
	}
}
