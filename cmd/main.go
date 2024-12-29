package main

import (
	"github.com/mohammed1146/skelton/internal/infrastructure/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	"github.com/mohammed1146/skelton/config"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/mohammed1146/skelton/docs"
	"github.com/mohammed1146/skelton/internal/initializer"
)

func main() {
	// Initialize the application
	app, err := initializer.InitializeApp()
	if err != nil {
		log.Fatalf("Initialization error: %v", err)
	}
	defer app.DB.Close()

	e := echo.New()

	// Swagger setup
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Handlers
	userHandler := handler.NewUserHandler(app.UserService)
	spacecraftHandler := handler.NewSpacecraftHandler(app.SpacecraftService)

	// Routes

	// Routes using Gorilla Mux
	r := mux.NewRouter()
	r.HandleFunc("/register", userHandler.Register).Methods(http.MethodPost)
	r.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)
	r.HandleFunc("/spacecrafts", spacecraftHandler.ListSpacecrafts).Methods(http.MethodGet)

	// Start http server.
	port := config.GetAppPort()
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
