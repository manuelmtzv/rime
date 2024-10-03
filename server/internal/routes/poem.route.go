package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterPoemRoutes(r *chi.Mux) {
	r.Group(func(r chi.Router) {
		r.Get("/poems", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Poems"))
		})
		r.Get("/poems/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Poem"))
		})
		r.Post("/poems", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Create Poem"))
		})
		r.Put("/poems/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Update Poem"))
		})
		r.Delete("/poems/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Delete Poem"))
		})
	})
}
