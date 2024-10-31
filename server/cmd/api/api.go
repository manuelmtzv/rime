package main

import (
	"net/http"
	"rime-api/internal/store"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type application struct {
	config config
	store  store.Storage
	logger *zap.SugaredLogger
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", app.register)
			r.Post("/login", app.login)
		})

		r.Route("/users", func(r chi.Router) {
			r.Get("/", app.findUsers)
			r.Get("/{id}", app.findOneUser)
		})

		r.Route("/writtings", func(r chi.Router) {
			r.Get("/", app.findWrittings)
			r.Get("/{id}", app.findOneWritting)
			r.Post("/", app.createWritting)
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 45,
		ReadTimeout:  time.Second * 20,
		IdleTimeout:  time.Minute,
	}

	app.logger.Infof("server has started at %s", app.config.addr)

	return server.ListenAndServe()
}
