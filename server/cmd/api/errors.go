package main

import (
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("internal error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusInternalServerError, app.getLocaleMessage(r, &i18n.LocalizeConfig{MessageID: "Error.InternalServerError"}))
}

func (app *application) forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	app.logger.Warnw("forbidden", "method", r.Method, "path", r.URL.Path, "error")

	writeJSONError(w, http.StatusForbidden, app.getLocaleMessage(r, &i18n.LocalizeConfig{MessageID: "Error.Forbidden"}))
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	msg := app.getLocaleMessage(r, &i18n.LocalizeConfig{MessageID: "Error.Conflict"})
	if err != nil {
		msg = err.Error()
	}

	app.logger.Errorf("conflict response", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusConflict, msg)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	msg := "not found"
	if err != nil {
		msg = err.Error()
	}

	app.logger.Warnf("Not Found: %s %s - %s", r.Method, r.URL.Path, msg)
	writeJSONError(w, http.StatusNotFound, msg)
}

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusUnauthorized, app.getLocaleMessage(r, &i18n.LocalizeConfig{MessageID: "Error.Unauthorized"}))
}

func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("unauthorized basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	writeJSONError(w, http.StatusUnauthorized, app.getLocaleMessage(r, &i18n.LocalizeConfig{MessageID: "Error.Unauthorized"}))
}

func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	app.logger.Warnw("rate limit exceeded", "method", r.Method, "path", r.URL.Path)
	w.Header().Set("Retry-After", retryAfter)

	msg := "rate limit exceeded, retry after: " + retryAfter

	localizer := app.getLocalizerFromContext(r)

	if localizer != nil {
		msg = localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Error.RateLimitExceeded", TemplateData: map[string]string{"RetryAfter": retryAfter}})
	}

	writeJSONError(w, http.StatusTooManyRequests, msg)
}
