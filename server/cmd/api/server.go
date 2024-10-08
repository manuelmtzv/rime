package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rime-server/internal/database"
	"rime-server/internal/models"
	"rime-server/internal/utils"
	"time"
)

type config struct {
	port int
	env  string
}

type application struct {
	logger *log.Logger
	models models.Models
	config config
}

type server struct {
	port int
}

func NewServer() *http.Server {
	port := utils.GetEnvAsIntOrThrow("PORT")
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := database.New()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	defer db.Close()

	logger.Printf("Connected to database")

	app := &application{
		logger: logger,
		models: models.NewModels(db),
	}

	newServer := &server{
		port: port,
	}

	logger.Printf("Starting %s server on %s", "API", fmt.Sprintf(":%d", newServer.port))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      app.RegisterRouter(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Failed to start server: %v", err)
	}

	return server
}
