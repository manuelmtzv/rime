package main

import (
	"errors"
	"net/http"
	"rime-api/internal/validator"

	"github.com/go-chi/chi/v5"
)

func (app *application) findUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.store.Users.FindAll(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.jsonResponse(w, http.StatusOK, users)

	if err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) findOneUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if !validator.IsValidUUID(id) {
		app.badRequestResponse(w, r, errors.New("user id must be a valid UUID"))
		return
	}

	user, err := app.store.Users.FindOne(r.Context(), id)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if user == nil {
		app.notFoundResponse(w, r, errors.New("user not found"))
		return
	}
}
