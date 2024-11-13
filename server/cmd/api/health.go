package main

import (
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	localizer := app.getLocalizerFromContext(r)

	localizeConfigWelcome := i18n.LocalizeConfig{
		MessageID: "Health.Ok",
	}
	message, _ := localizer.Localize(&localizeConfigWelcome)

	app.logger.Info(message)

	if err := app.jsonResponse(w, http.StatusOK, map[string]interface{}{"message": message}); err != nil {
		app.internalServerError(w, r, err)
	}
}
