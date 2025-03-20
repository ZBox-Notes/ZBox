package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/ZBox-Notes/ZBox/backend/api/boxes"
	"github.com/ZBox-Notes/ZBox/backend/api/notes"
	notesboxes "github.com/ZBox-Notes/ZBox/backend/api/notes_boxes"
	"github.com/ZBox-Notes/ZBox/backend/api/users"
	database "github.com/ZBox-Notes/ZBox/backend/db/postgres"
	middleware "github.com/ZBox-Notes/ZBox/backend/middleware"
	model "github.com/ZBox-Notes/ZBox/backend/models/generated_model"

	"github.com/gorilla/mux"
)

func main() {
	greet()
	slog.Info("Starting server...")

	// Connect to the database
	slog.Info("Connecting to the database...")
	conn, connErr := database.NewPostgres()
	if connErr != nil {
		slog.Error("Failed to connect to the database")
		panic(connErr)
	}
	defer conn.DB.Close(context.Background())

	// Instantiate services
	slog.Info("Instantiating services...")
	queries := model.New(conn.DB)
	userService := users.NewService(queries)
	notesService := notes.NewService(queries)
	boxesService := boxes.NewService(queries)
	notesboxesService := notesboxes.NewService(queries)

	// Register handlers
	slog.Info("Registering handlers...")
	r := mux.NewRouter()
	userService.RegisterHandlers(r)
	notesService.RegisterHandlers(r)
	boxesService.RegisterHandlers(r)
	notesboxesService.RegisterHandlers(r)

	// Add middleware
	slog.Info("Adding middleware...")
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.AuthMiddleware)

	// Start the server
	slog.Info("Server started successfully")
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func greet() {
	welcomeMsg := `    

  _____  ______   ______  ______         ____    _____       _____ 
 /    / /     /|  \     \|\     \    ____\_  \__ \    \     /    / 
|     |/     / |   |     |\|     |  /     /     \ \    |   |    /  
|\____\\    / /    |     |/____ /  /     /\      | \    \ /    /   
 \|___|/   / /     |     |\     \ |     |  |     |  \    |    /    
    /     /_/____  |     | |     ||     |  |     |  /    |    \    
   /     /\      | |     | |     ||     | /     /| /    /|\    \   
  /_____/ /_____/|/_____/|/_____/||\     \_____/ ||____|/ \|____|  
  |    |/|     | ||    |||     | || \_____\   | / |    |   |    |  
  |____| |_____|/ |____|/|_____|/  \ |    |___|/  |____|   |____|  
                                    \|____|     

Star ZBox on Github: https://github.com/ZBox-Notes/ZBox
##################################################################
`

	slog.Info(welcomeMsg)
}
