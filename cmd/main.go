package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // <-- Add this line to register the Postgres driver

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals"
	approvalStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/communication"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/config"
	cron2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/cron"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/events"
	eventsStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/events/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/router"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media"
	mediaStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media/youtube"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry"
	ministryStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach"
	outreachStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles"
	rolesStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	userStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/storage"
	commonDb "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/db"
	commonlogger "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/logger"
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

	logger := commonlogger.New(cfg)

	// Connect to the PostgreSQL database using sqlx.
	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		logger.Sugar().Fatalf("failed to connect to database: %v", err)
	}

	if err := commonDb.RunMigrations(db.DB); err != nil {
		logger.Sugar().Fatalf("failed to run migrations: %v", err)
	}

	rolesRepo := rolesStorage.NewRolesRepository(db)
	rolesService := roles.NewRoleService(logger, rolesRepo)

	communicationService := communication.NewCommunicationService(
		cfg.TwilioAccountSID,
		cfg.TwilioAuthToken,
		cfg.TwilioNumber,
	)

	approvalRepository := approvalStorage.NewApprovalRepository(db)
	approvalService := approvals.NewApprovalService(logger, approvalRepository, nil, rolesService)

	// Declare ministryService early so we can pass it into userService
	var ministryService ministry.MinistryService

	userRepo := userStorage.NewUserRepository(db)
	userService := user.NewUserService(logger, userRepo, cfg.JwtSecret, nil, approvalService, rolesService)

	ministryRepo := ministryStorage.NewMinistryRepository(db)
	ministryService = ministry.NewMinistryService(
		logger,
		ministryRepo,
		communicationService,
		userService, // uses interface, no cycle
		approvalService,
	)

	userService.SetMinistryService(ministryService)
	approvalService.SetUserService(userService)

	outreachRepo := outreachStorage.NewOutreachRepository(db)
	outreachService := outreach.NewOutreachService(logger, outreachRepo)

	mediaRepo := mediaStorage.NewMediaRepository(db)
	mediaService := media.NewMediaService(logger, mediaRepo)

	eventsRepo := eventsStorage.NewEventsRepository(db)
	eventsService := events.NewEventsService(logger, eventsRepo)

	ytClient := youtube.NewYouTubeClient(cfg.YoutubeAPIKey, cfg.YoutubeChannelID)
	mediaCronJob := cron2.NewMediaCronJob(logger, ytClient, mediaService)

	mediaCronJob.RunOnce()

	mediaCronJob.Start()

	mux := router.New(
		logger,
		cfg.JwtSecret,
		userService,
		ministryService,
		outreachService,
		mediaService,
		approvalService,
		eventsService,
	)

	logger.Sugar().Infof("Server starting on port %s", cfg.Port)

	addr := fmt.Sprintf(":%s", cfg.Port)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
