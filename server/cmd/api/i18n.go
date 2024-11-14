package main

import (
	"encoding/json"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type localizerKey string

const localizerCtx localizerKey = "localizer"

func initI18n() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bundle.MustLoadMessageFile("i18n/active.en.json")
	bundle.MustLoadMessageFile("i18n/active.es.json")

	return bundle
}

func (app *application) getLocalizerFromContext(r *http.Request) *i18n.Localizer {
	localizer, _ := r.Context().Value(localizerCtx).(*i18n.Localizer)

	if localizer == nil {
		localizer = i18n.NewLocalizer(app.i18n.bundle, r.Header.Get("Accept-Language"))
	}

	return localizer
}

func (app *application) getLocaleMessage(r *http.Request, cfg *i18n.LocalizeConfig) string {
	localizer := app.getLocalizerFromContext(r)
	return localizer.MustLocalize(cfg)
}
