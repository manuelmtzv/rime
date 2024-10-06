package server

import (
	"net/http"
	"rime-server/internal/routes"

	"github.com/go-chi/chi/v5"
)

func RegisterRouter() http.Handler {
	r := chi.NewRouter()

	routes.RegisterPoemRoutes(r)

	return r
}
