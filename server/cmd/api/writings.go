package main

import (
	"errors"
	"net/http"
	"rime-api/internal/models"

	"github.com/go-chi/chi/v5"
)

type CreateWritingPayload struct {
	Type    string `json:"type" validate:"oneof=poetry"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (app *application) createWriting(w http.ResponseWriter, r *http.Request) {
	var payload CreateWritingPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := app.getUserFromContext(r)

	writting := &models.Writing{
		Type:     payload.Type,
		Title:    payload.Title,
		Content:  payload.Content,
		AuthorID: user.ID,
	}

	if err := app.store.Writings.Create(r.Context(), writting); err != nil {
		app.internalServerErrorBasic(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, writting); err != nil {
		app.internalServerErrorBasic(w, r, err)
	}
}

func (app *application) findWritings(w http.ResponseWriter, r *http.Request) {
	writings, err := app.store.Writings.FindAll(r.Context())
	if err != nil {
		app.internalServerErrorBasic(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, writings); err != nil {
		app.internalServerErrorBasic(w, r, err)
	}
}

func (app *application) findOneWriting(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	writting, err := app.store.Writings.FindOne(r.Context(), id)
	if err != nil {
		app.internalServerErrorBasic(w, r, err)
		return
	}

	if writting == nil {
		app.notFoundResponse(w, r, errors.New("writting not found"))
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, writting); err != nil {
		app.internalServerErrorBasic(w, r, err)
	}
}

func (app *application) findWritingDetails(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	writtingDetails, err := app.store.Writings.FindDetails(r.Context(), id)
	if err != nil {
		app.internalServerErrorBasic(w, r, err)
		return
	}

	if writtingDetails == nil {
		app.notFoundResponse(w, r, errors.New("writing not found"))
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, writtingDetails); err != nil {
		app.internalServerErrorBasic(w, r, err)
	}
}
