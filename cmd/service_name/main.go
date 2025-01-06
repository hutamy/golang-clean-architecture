package main

import (
	"os"

	"github.com/hutamy/golang-clean-architecture/config"
	"github.com/hutamy/golang-clean-architecture/internal/delivery/http"
	"github.com/hutamy/golang-clean-architecture/internal/delivery/http/handlers"
	"github.com/hutamy/golang-clean-architecture/internal/repository/cache"
	"github.com/hutamy/golang-clean-architecture/internal/repository/mongo"
	"github.com/hutamy/golang-clean-architecture/internal/repository/postgres"
	"github.com/hutamy/golang-clean-architecture/internal/service"
	"github.com/hutamy/golang-clean-architecture/pkg/logger"
	"github.com/hutamy/golang-clean-architecture/pkg/middleware"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig() // Load from environment or config file

	// Logger configuration
	logConfig := logger.Config{
		Level:      "debug",
		JSONFormat: true,
	}

	// Initialize logger
	log := logger.New(logConfig)
	defer log.Sync()

	// Initialize PostgresDB
	pgDbManager := postgres.NewDBManager(cfg.Database.MasterDSN, cfg.Database.Replicas)

	// Initialize MongoDB
	mongoDbManager := mongo.NewDBManager(cfg.Mongo.URI, cfg.Mongo.Database)
	defer mongoDbManager.Close()

	// Initialize Redis connection
	redisManager := cache.NewDBManager(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)
	defer redisManager.Close()

	// Initialize the repository
	userRepo := postgres.NewUserRepository(pgDbManager)
	logsRepo := mongo.NewLogsRepository(mongoDbManager)
	cacheRepo := cache.NewCacheRepository(redisManager)

	userService := service.NewUserService(
		userRepo,
		logsRepo,
		cacheRepo,
		log,
	)
	authService := service.NewAuthService(
		userRepo,
		cfg,
		log,
	)

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	e := echo.New()

	// Logger middleware
	e.Use(m.Logger())
	e.Use(m.Recover())

	// Secret key for JWT
	secretKey := cfg.Jwt.SecretKey

	// Initialize AuthMiddleware
	authMiddleware := middleware.NewAuthMiddleware(secretKey)

	http.RegisterRoutes(e, authMiddleware, userHandler, authHandler)

	log.Info("Starting the application")
	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		log.Fatal("server failed to start: %v", err)
	}
}
