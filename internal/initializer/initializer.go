package initializer

import (
	"context"
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/mohammed1146/skelton/internal/repository"
	"github.com/mohammed1146/skelton/internal/service"
	"log"

	"github.com/mohammed1146/skelton/config"
)

type App struct {
	DB                *sql.DB
	SpacecraftService service.SpacecraftService
	UserService       service.UserService
}

func InitializeApp() (*App, error) {
	ctx := context.Background()

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, relying on system environment variables")
	}

	// Initialize DB.
	dbConfig := config.GetDBConfig()
	db, err := config.NewDB(ctx, dbConfig)
	if err != nil {
		return nil, err
	}
	log.Println("Database connection established.")

	// Run DB migrations.
	config.RunMigrations(dbConfig)

	// Initialize spacecraft module.
	spacecraftRepo := repository.NewSpacecraftRepository(db)
	spacecraftService := service.NewSpacecraftService(spacecraftRepo)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	return &App{
		DB:                db,
		SpacecraftService: spacecraftService,
		UserService:       userService,
	}, nil
}
