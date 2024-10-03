package server

import (
	"net/http"
)

func (s *Server) RegisterRouter() http.Handler {
	router := http.NewServeMux()

	return router
}
