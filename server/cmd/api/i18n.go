package main

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func initI18n() (*i18n.Bundle, *i18n.Localizer) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bundle.LoadMessageFile("resources/i18n/en.json")
	bundle.LoadMessageFile("resources/i18n/es.json")

	localizer := i18n.NewLocalizer(bundle, language.English.String(), language.Spanish.String())

	return bundle, localizer
}
