package main

import (
	"errors"
	"net/http"
	"rime-api/internal/hash"
	"rime-api/internal/models"
	"rime-api/internal/validations"

	"github.com/go-chi/chi/v5"
)

type CreateUserPayload struct {
	Name     string `json:"name" validate:"required,max=100"`
	LastName string `json:"last_name" validate:"required,max=100"`
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, errors.New("the provided JSON payload is invalid"))
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if user, err := app.store.Users.FindByEmail(r.Context(), payload.Email); err == nil && user != nil {
		app.badRequestResponse(w, r, errors.New("email already exists"))
		return
	}

	hashedPassword, err := hash.HashPassword(payload.Password)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	user := &models.User{
		Name:     payload.Name,
		LastName: payload.LastName,
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashedPassword,
	}

	if err := app.store.Users.Create(r.Context(), user); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	userResponse := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	if err := app.jsonResponse(w, http.StatusCreated, userResponse); err != nil {
		app.internalServerError(w, r, err)
	}
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
