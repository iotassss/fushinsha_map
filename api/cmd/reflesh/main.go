package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/iotassss/fushinsha-map-api/internal/repository/gormrepo"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// load .env
	_ = godotenv.Load(".env")
	env := os.Getenv("APP_ENV")
	if env == "" {
		slog.Error(" environment variables", slog.Any("error", "APP_ENV"))
		return
	}
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		slog.Error("Missing required environment variables", slog.Any("error", "APP_PORT"))
		return
	}
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	if mysqlDatabase == "" {
		slog.Error("Missing required environment variables", slog.Any("error", "MYSQL_DATABASE"))
		return
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		slog.Error("Missing required environment variables", slog.Any("error", "MYSQL_USER"))
		return
	}
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword == "" {
		slog.Error("Missing required environment variables", slog.Any("error", "MYSQL_PASSWORD"))
		return
	}
	// dbHost := os.Getenv("DB_HOST")
	// if dbHost == "" {
	// 	slog.Error("Missing required environment variables", slog.Any("error", "DB_HOST"))
	// 	return
	// }
	dbHost := "localhost"
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	if googleClientID == "" {
		slog.Error("Missing required environment variables", slog.Any("error", "GOOGLE_CLIENT_ID"))
		return
	}

	// logger
	logWriter := os.Stdout
	slogJSONHandler := slog.NewJSONHandler(logWriter, nil)
	logger := slog.New(slogJSONHandler)
	slog.SetDefault(logger)

	// database
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPassword, dbHost, mysqlDatabase)
	if env == "development" {
		slog.Info("connecting to database", slog.Any("dsn", dbDSN))
	}
	db, err := gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect to database", slog.Any("error", err))
		return
	}
	err = db.AutoMigrate(
		&gormrepo.PersonModel{},
		&gormrepo.UserModel{},
	)
	if err != nil {
		slog.Error("failed to migrate database", slog.Any("error", err))
		return
	}

	ctx := context.Background()
	personRepo := gormrepo.NewPersonRepository(db)
	if err := personRepo.ResetTable(ctx); err != nil {
		slog.Error("failed to reset existing data during dummy data seeding", slog.Any("error", err))
		return
	}
}
