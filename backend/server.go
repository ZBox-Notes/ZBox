package main

import (
	"backend/api/users"
	database "backend/db"
	model "backend/models/generated_model"
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	conn, connErr := database.NewPostgres()
	if connErr != nil {
		panic(connErr)
	}
	defer conn.DB.Close(context.Background())

	// Instantiate services
	queries := model.New(conn.DB)
	userService := users.NewService(queries)

	// Register handlers
	r := mux.NewRouter()
	userService.RegisterHandlers(r)

	// Add middleware
	r.Use(loggingMiddleware)

	// Start the server
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
