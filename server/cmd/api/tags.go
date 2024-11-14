package main

import (
	"errors"
	"net/http"
	"rime-api/internal/models"
)

type CreateTagPayload struct {
	Name string `json:"name" validate:"required,min=3,max=50"`
}

func (app *application) createTag(w http.ResponseWriter, r *http.Request) {
	var payload CreateTagPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	existingTag, _ := app.store.Tags.FindOneByName(r.Context(), payload.Name)
	if existingTag != nil {
		app.conflictResponse(w, r, errors.New("tag already exists"))
		return
	}

	tag := &models.Tag{
		Name: payload.Name,
	}

	if err := app.store.Tags.Create(r.Context(), tag); err != nil {
		app.internalServerErrorBasic(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, tag); err != nil {
		app.internalServerErrorBasic(w, r, err)
	}

}

func (app *application) findTags(w http.ResponseWriter, r *http.Request) {
	tags, err := app.store.Tags.FindAll(r.Context())
	if err != nil {
		app.internalServerErrorBasic(w, r, err)
		return
	}

	err = app.jsonResponse(w, http.StatusOK, tags)

	if err != nil {
		app.internalServerErrorBasic(w, r, err)
	}
}

func (app *application) findPopularTags(w http.ResponseWriter, r *http.Request) {
	tags, err := app.store.Tags.FindPopular(r.Context())
	if err != nil {
		app.internalServerErrorBasic(w, r, err)
		return
	}

	err = app.jsonResponse(w, http.StatusOK, tags)

	if err != nil {
		app.internalServerErrorBasic(w, r, err)
	}
}
