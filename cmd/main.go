package main

import (
	"fmt"
	commonDb "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/db"
	logger2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/logger"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // <-- Add this line to register the Postgres driver

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/config"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/router"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry"
	ministryStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles"
	rolesStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	userStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/storage"
)

func main() {
	// Load environment variables from .env file (only in development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	logger := logger2.New(cfg)

	// Connect to the PostgreSQL database using sqlx.
	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		logger.Fatalf("failed to connect to database: %v", err)
	}

	if err := commonDb.RunMigrations(db.DB); err != nil {
		logger.Fatalf("failed to run migrations: %v", err)
	}

	userRepo := userStorage.NewUserRepository(db)
	userService := user.NewUserService(logger, userRepo, cfg.JwtSecret)

	rolesRepo := rolesStorage.NewRolesRepository(db)
	rolesService := roles.NewRoleService(logger, rolesRepo)

	ministryRepo := ministryStorage.NewMinistryRepository(db)
	ministryService := ministry.NewMinistryService(logger, ministryRepo, cfg.JwtSecret)

	mux := router.New(logger, userService, rolesService, ministryService, cfg.JwtSecret)

	logger.Infof("Server starting on port %s", cfg.Port)

	addr := fmt.Sprintf(":%s", cfg.Port)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
