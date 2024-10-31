package main

import (
	"errors"
	"net/http"
	"rime-api/internal/models"

	"github.com/go-chi/chi/v5"
)

type CreateWrittingPayload struct {
	Type    string `json:"type" validate:"oneof=poem"`
	Content string `json:"content" validate:"required"`
}

type WrittingResponse struct {
}

func (app *application) createWritting(w http.ResponseWriter, r *http.Request) {
	var payload CreateWrittingPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	writting := &models.Writting{
		Type:    payload.Type,
		Content: payload.Content,
	}

	if err := app.store.Writtings.Create(r.Context(), writting); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, writting); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) findWrittings(w http.ResponseWriter, r *http.Request) {
	writtings, err := app.store.Writtings.FindAll(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, writtings); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) findOneWritting(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	writting, err := app.store.Writtings.FindOne(r.Context(), id)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if writting == nil {
		app.notFoundResponse(w, r, errors.New("writting not found"))
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, writting); err != nil {
		app.internalServerError(w, r, err)
	}
}
