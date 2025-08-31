package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	dblHost := os.Getenv("DB_HOST")
	if dblHost == "" {
		slog.Error("Missing required environment variables", slog.Any("error", "DB_HOST"))
		return
	}

	// logger
	logWriter := os.Stdout
	slogJSONHandler := slog.NewJSONHandler(logWriter, nil)
	logger := slog.New(slogJSONHandler)
	slog.SetDefault(logger)

	// database
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPassword, dblHost, mysqlDatabase)
	if env == "development" {
		slog.Info("connecting to database", slog.Any("dsn", dbDSN))
	}
	db, err := gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect to database", slog.Any("error", err))
		return
	}
	err = db.AutoMigrate(
		&gormrepo.CityModel{},
	)
	if err != nil {
		slog.Error("failed to migrate database", slog.Any("error", err))
		return
	}

	// dummy data
	if env == "development" {
		cityRepo := gormrepo.NewCityRepository(db, context.Background())
		if err := cityRepo.ResetTable(); err != nil {
			slog.Error("failed to reset existing data during dummy data seeding", slog.Any("error", err))
			return
		}
		if err := cityRepo.SeedDummyCity(); err != nil {
			slog.Error("failed to seed dummy data", slog.Any("error", err))
			return
		}
	}

	// // handler
	// loginHandler := handler.NewLoginHandler(db)

	// router
	r := gin.Default()

	// CORSミドルウェアを追加
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 認証不要なAPI
	api := r.Group("/api")
	{
		// api.POST("/login", loginHandler)
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
		api.GET("/cities", func(c *gin.Context) {
			cityRepo := gormrepo.NewCityRepository(db, context.Background())
			cities, err := cityRepo.FindAll()
			if err != nil {
				slog.Error("failed to get cities", slog.Any("error", err))
				c.JSON(500, gin.H{
					"error": "failed to get cities",
				})
				return
			}
			result := make([]map[string]string, 0, len(cities))
			for _, city := range cities {
				result = append(result, map[string]string{
					"id":   city.ID,
					"name": city.Name,
				})
			}
			c.JSON(200, gin.H{
				"cities": result,
			})
		})
	}

	// // 認証が必要なAPI
	// authorized := r.Group("/api")
	// authorized.Use(middleware.AuthMiddleware())
	// {
	// }

	r.Run() // デフォルトで :8080 で起動
}
