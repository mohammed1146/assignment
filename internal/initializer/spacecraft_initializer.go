package initializer

import (
	"context"
	"database/sql"

	"github.com/mohammed1146/skelton/internal/infrastructure/handler"
	"github.com/mohammed1146/skelton/internal/repository"
	"github.com/mohammed1146/skelton/internal/service"
)

func InitializeSpacecraftModule(ctx context.Context, db *sql.DB) *handler.SpacecraftHandler {
	repo := repository.NewSpacecraftRepository(db)
	service := service.NewSpacecraftService(repo)
	return handler.NewSpacecraftHandler(service)
}
