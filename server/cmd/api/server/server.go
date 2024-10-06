package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"rime-server/internal/database"
	"rime-server/internal/utils"
	"time"
)

type server struct {
	port int
	db   *sql.DB
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

	newServer := &server{
		port: port,
		db:   db,
	}

	logger.Printf("Starting %s server on %s", "API", fmt.Sprintf(":%d", newServer.port))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      RegisterRouter(),
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
