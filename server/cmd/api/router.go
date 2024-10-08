package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) RegisterRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/healthcheck", app.healthcheckHandler)

	r.Group(func(r chi.Router) {
		r.Get("/users", app.GetUsers)
		r.Get("/users/{id}", app.GetUser)
		r.Post("/users", app.CreateUser)
		r.Put("/users/{id}", app.UpdateUser)
		r.Delete("/users/{id}", app.DeleteUser)
	})

	r.Group(func(r chi.Router) {
		r.Get("/poems", app.GetPoems)
		r.Get("/poems/{id}", app.GetPoem)
		r.Post("/poems", app.CreatePoem)
		r.Put("/poems/{id}", app.UpdatePoem)
		r.Delete("/poems/{id}", app.DeletePoem)
	})

	return r
}
