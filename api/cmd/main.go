package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iotassss/fushinsha-map-api/internal/auth"
	handler "github.com/iotassss/fushinsha-map-api/internal/handler/api"
	"github.com/iotassss/fushinsha-map-api/internal/middleware"
	"github.com/iotassss/fushinsha-map-api/internal/repository/gormrepo"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
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
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		slog.Error("Missing required environment variables", slog.Any("error", "DB_HOST"))
		return
	}
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

	// repository
	userRepo := gormrepo.NewUserRepository(db)
	personRepo := gormrepo.NewPersonRepository(db)

	// auth verifier
	googleAuthVerifier := auth.NewGoogleAuthVerifier(googleClientID)

	// usecase
	findAndCreateUserByIDTokenInteractor := usecase.NewFindOrCreateUserByIDTokenInteractor(
		userRepo,
		googleAuthVerifier,
	)
	getPersonsInteractor := usecase.NewGetPersonsInteractor(personRepo)
	getPersonDetailInteractor := usecase.NewGetPersonDetailInteractor(personRepo)
	createPersonInteractor := usecase.NewCreatePersonInteractor(personRepo)
	updatePersonInteractor := usecase.NewUpdatePersonInteractor(personRepo)

	// middleware
	authMiddleware := middleware.NewAuthMiddleware(findAndCreateUserByIDTokenInteractor)

	// handler
	getPersonsHandler := handler.NewGetPersonsHandler(getPersonsInteractor)
	getPersonDetailHandler := handler.NewGetPersonDetailHandler(getPersonDetailInteractor)
	createPersonHandler := handler.NewCreatePersonHandler(createPersonInteractor)
	updatePersonHandler := handler.NewUpdatePersonHandler(updatePersonInteractor)

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

	// APIグループ
	api := r.Group("/api")

	// 認証不要
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	api.GET("/persons", getPersonsHandler.Handle)
	api.GET("/persons/:uuid", getPersonDetailHandler.Handle)

	// 認証が必要なAPI
	authorized := api.Group("")
	// TODO: 認証機能を適切に実装して以下の処理を有効化する
	_ = authMiddleware
	// authorized.Use(authMiddleware.Auth())
	{
		authorized.GET("/me")
		authorized.POST("/persons", createPersonHandler.Handle)
		authorized.PUT("/persons/:uuid", updatePersonHandler.Handle)
	}

	r.Run() // デフォルトで :8080 で起動
}
