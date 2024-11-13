package main

import (
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	lang := r.FormValue("lang")
	accept := r.Header.Get("Accept-Language")

	localizeConfigWelcome := i18n.LocalizeConfig{
		MessageID: "health.ok",
	}
	message, _ := app.i18n.localizer.Localize(&localizeConfigWelcome)

	w.Write([]byte(message))
}
