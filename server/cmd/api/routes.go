package main

import (
	"rime-api/internal/env"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(app.LocalizerMiddleware)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{env.GetString("CORS_ALLOWED_ORIGIN", "http://localhost:5174")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		r.Route("/auth", func(r chi.Router) {
			r.With(app.AuthMiddleware).Get("/validate", app.validate)
			r.Get("/refresh", app.refreshToken)
			r.Post("/register", app.register)
			r.Post("/login", app.login)
		})

		r.Route("/users", func(r chi.Router) {
			r.Get("/", app.findUsers)
			r.Get("/{id}", app.findOneUser)
			r.Get("/popular", app.findPopular)
		})

		r.Route("/writings", func(r chi.Router) {
			r.Get("/", app.findWritings)
			r.Get("/{id}", app.findOneWriting)
			r.Get("/{id}/details", app.findWritingDetails)
			r.With(app.AuthMiddleware).Post("/", app.createWriting)
		})

		r.Route("/tags", func(r chi.Router) {
			r.Get("/", app.findTags)
			r.Get("/popular", app.findPopularTags)
			r.With(app.AuthMiddleware).Post("/", app.createTag)
		})

		r.Route("/likes", func(r chi.Router) {
			r.With(app.AuthMiddleware).Post("/writings/{id}", app.likeWriting)
			r.With(app.AuthMiddleware).Delete("/writings/{id}", app.unlikeWriting)
			r.With(app.AuthMiddleware).Post("/comments/{id}", app.likeComment)
			r.With(app.AuthMiddleware).Delete("/comments/{id}", app.unlikeComment)
		})
	})

	app.router = r
}
