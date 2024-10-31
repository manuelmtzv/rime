package main

import (
	"errors"
	"net/http"
	"rime-api/internal/hash"
	"rime-api/internal/models"
)

type RegisterPayload struct {
	Name     string `json:"name" validate:"required,max=100"`
	Lastname string `json:"lastname" validate:"required,max=100"`
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginPayload struct {
	Identifier string `json:"identifier" validator:"required"`
	Password   string `json:"password" validator:"required"`
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	var payload RegisterPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if user, err := app.store.Users.FindByIdentifier(r.Context(), payload.Email); err == nil && user != nil {
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
		Lastname: payload.Lastname,
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
		Lastname:  user.Lastname,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	if err := app.jsonResponse(w, http.StatusCreated, userResponse); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {

}
