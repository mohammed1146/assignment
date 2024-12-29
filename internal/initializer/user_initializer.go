package initializer

import (
	"context"
	"database/sql"

	"github.com/mohammed1146/skelton/internal/infrastructure/handler"
	"github.com/mohammed1146/skelton/internal/repository"
	"github.com/mohammed1146/skelton/internal/service"
)

func InitializeUserModule(ctx context.Context, db *sql.DB) *handler.UserHandler {
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	return handler.NewUserHandler(service)
}
