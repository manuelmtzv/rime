package routes

import (
	"rime-server/internal/controllers"

	"github.com/go-chi/chi/v5"
)

func RegisterPoemRoutes(r *chi.Mux) {
	poemController := &controllers.PoemController{}

	r.Group(func(r chi.Router) {
		r.Get("/poems", poemController.GetPoems)
		r.Get("/poems/{id}", poemController.GetPoem)
		r.Post("/poems", poemController.CreatePoem)
		r.Put("/poems/{id}", poemController.UpdatePoem)
		r.Delete("/poems/{id}", poemController.DeletePoem)
	})
}
