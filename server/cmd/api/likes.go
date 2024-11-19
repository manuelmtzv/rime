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

	like := &models.WritingLike{
		AuthorID:  user.ID,
		WritingID: writing.ID,
	}

	if err := app.store.Likes.CreateWritingLike(r.Context(), like); err != nil {
		app.internalServerErrorBasic(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, like); err != nil {
		app.internalServerErrorBasic(w, r, err)
	}
}

func (app *application) unlikeWriting(w http.ResponseWriter, r *http.Request) {
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

	if err := app.store.Likes.DeleteWritingLike(r.Context(), user.ID, writing.ID); err != nil {
		app.internalServerErrorBasic(w, r, err)
	}

	if err := app.emptyResponse(w, http.StatusNoContent); err != nil {
		app.internalServerErrorBasic(w, r, err)
	}
}

func (app *application) likeComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	comment, err := app.store.Comments.FindOne(r.Context(), id)
	if err != nil {
		app.internalServerErrorBasic(w, r, err)
		return
	}

	if comment == nil {
		app.notFoundResponse(w, r, app.getLocaleError(r, constants.CommentNotFound, nil))
		return
	}

	user := app.getUserFromContext(r)

	like := &models.CommentLike{
		AuthorID:  user.ID,
		CommentID: comment.ID,
	}

	if err := app.store.Likes.CreateCommentLike(r.Context(), like); err != nil {
		app.internalServerErrorBasic(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, like); err != nil {
		app.internalServerErrorBasic(w, r, err)
	}
}

func (app *application) unlikeComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	comment, err := app.store.Comments.FindOne(r.Context(), id)
	if err != nil {
		app.internalServerErrorBasic(w, r, err)
		return
	}

	if comment == nil {
		app.notFoundResponse(w, r, app.getLocaleError(r, constants.CommentNotFound, nil))
		return
	}

	user := app.getUserFromContext(r)

	if err = app.store.Likes.DeleteCommentLike(r.Context(), user.ID, comment.ID); err != nil {
		app.internalServerErrorBasic(w, r, err)
	}

	if err := app.emptyResponse(w, http.StatusNoContent); err != nil {
		app.internalServerErrorBasic(w, r, err)
	}
}
