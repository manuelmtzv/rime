package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"rime-server/internal/database"
	"strconv"
	"time"
)

type Server struct {
	port int
	db   *sql.DB
}

func NewServer() *http.Server {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		panic(fmt.Sprintf("Invalid PORT: %v", err))
	}

	db, err := database.New()

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	newServer := &Server{
		port: port,
		db:   db,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRouter(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
