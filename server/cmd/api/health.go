package main

import (
	"net/http"
	"rime-api/internal/constants"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	msg := app.getMessageOrDefault(r, constants.HealthOk, "Ok", nil)

	app.logger.Info(msg)

	if err := app.jsonResponse(w, http.StatusOK, map[string]interface{}{"message": msg}); err != nil {
		app.internalServerError(w, r, err)
	}
}
