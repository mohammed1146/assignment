package router

import (
	"github.com/mohammed1146/skelton/internal/infrastructure/handler"

	"github.com/labstack/echo/v4"
)

// AppDependencies defines the dependencies required by the router
type AppDependencies struct {
	UserHandler       *handler.UserHandler
	SpacecraftHandler *handler.SpacecraftHandler
}

// NewRouter sets up the application routes
func NewRouter(app *AppDependencies) *echo.Echo {
	e := echo.New()

	// User Routes
	e.POST("/register", app.UserHandler.Register)
	e.POST("/login", app.UserHandler.Login)
	e.GET("/spacecrafts", app.SpacecraftHandler.ListSpacecrafts)

	// Swagger Route
	e.GET("/swagger/*", echo.WrapHandler)

	return e
}
