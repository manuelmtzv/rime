package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		panic(fmt.Sprintf("Invalid PORT: %v", err))
	}

	newServer := &Server{
		port: port,
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
