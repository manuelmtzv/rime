package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type localizerKey string

const (
	localizerCtx      localizerKey = "localizer"
	ErrorUserNotFound              = "Error.UserNotFound"
)

func initI18n() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bundle.MustLoadMessageFile("i18n/active.en.json")
	bundle.MustLoadMessageFile("i18n/active.es.json")

	return bundle
}

func (app *application) getLocalizerFromContext(r *http.Request) (*i18n.Localizer, error) {
	localizer, ok := r.Context().Value(localizerCtx).(*i18n.Localizer)
	if !ok {
		return nil, errors.New("localizer not found in context")
	}
	return localizer, nil
}

func (app *application) getLocaleMessage(r *http.Request, messageID string, templateData map[string]interface{}) (string, error) {
	localizer, err := app.getLocalizerFromContext(r)
	if err != nil {
		return "", err
	}
	return localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})
}

func (app *application) getMessageOrDefault(r *http.Request, messageID, defaultMessage string, templateData map[string]interface{}) string {
	msg, err := app.getLocaleMessage(r, messageID, templateData)
	if err != nil {
		return defaultMessage
	}
	return msg
}

func (app *application) getLocaleError(r *http.Request, messageID string, templateData map[string]interface{}) error {
	msg, err := app.getLocaleMessage(r, messageID, templateData)
	if err != nil {
		return err
	}
	return errors.New(msg)
}
