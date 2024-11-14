package main

import (
	"net/http"
	"rime-api/internal/constants"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error, messageID string) {
	app.logger.Errorw("internal error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	msg := app.getMessageOrDefault(r, messageID, "internal server error", nil)

	writeJSONError(w, http.StatusInternalServerError, msg)
}

func (app *application) internalServerErrorBasic(w http.ResponseWriter, r *http.Request, err error) {
	app.internalServerError(w, r, err, constants.ErrorInternalServerError)
}

func (app *application) forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	app.logger.Warnw("forbidden", "method", r.Method, "path", r.URL.Path)

	writeJSONError(w, http.StatusForbidden, app.getMessageOrDefault(r, constants.ErrorForbidden, "this action is forbidden", nil))
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	msg, localizeErr := app.getLocaleMessage(r, constants.ErrorConflict, nil)
	if err != nil && localizeErr == nil {
		msg = err.Error()
	}

	app.logger.Errorw("conflict response", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusConflict, msg)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	msg := "not found"
	if err != nil {
		msg = err.Error()
	}

	app.logger.Warnw("not found", "method", r.Method, "path", r.URL.Path, "error", msg)

	writeJSONError(w, http.StatusNotFound, msg)
}

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("unauthorized", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusUnauthorized, app.getMessageOrDefault(r, constants.ErrorUnauthorized, "unauthorized", nil))
}

func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("unauthorized basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	writeJSONError(w, http.StatusUnauthorized, app.getMessageOrDefault(r, constants.ErrorUnauthorized, "unauthorized", nil))
}

func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	app.logger.Warnw("rate limit exceeded", "method", r.Method, "path", r.URL.Path)

	w.Header().Set("Retry-After", retryAfter)

	msg := "rate limit exceeded, retry after: " + retryAfter
	if localizer, err := app.getLocalizerFromContext(r); err == nil {
		msg, _ = localizer.Localize(&i18n.LocalizeConfig{
			MessageID:    constants.ErrorRateLimitExceeded,
			TemplateData: map[string]string{"RetryAfter": retryAfter},
		})
	}

	writeJSONError(w, http.StatusTooManyRequests, msg)
}
