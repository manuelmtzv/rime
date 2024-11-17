package main

import (
	"net/http"
	"rime-api/internal/constants"
	"rime-api/internal/models"

	"github.com/go-chi/chi/v5"
)

func (app *application) likeWriting(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	writing, err := app.store.Writings.FindOne(r.Context(), id)
	if err != nil {
		app.internalServerErrorBasic(w, r, err)
		return
	}

	if writing == nil {
		app.notFoundResponse(w, r, app.getLocaleError(r, constants.WritingNotFound, nil))
		return
	}

	user := app.getUserFromContext(r)

	like := &models.Like{
		UserID: user.ID,
	}

	if err := app.store.Likes.LikeWriting(r.Context(), like, writing.ID); err != nil {
		app.internalServerErrorBasic(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, like); err != nil {
		app.internalServerErrorBasic(w, r, err)
	}
}

func (app *application) likeComment(w http.ResponseWriter, r *http.Request) {

}
