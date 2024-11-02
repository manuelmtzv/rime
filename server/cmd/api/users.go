package main

import (
	"errors"
	"net/http"
	"rime-api/internal/validations"

	"github.com/go-chi/chi/v5"
)

type userKey string

const userCtx userKey = "user"

type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

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

	if !validations.IsValidUUID(id) {
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

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}
}
