package main

import (
	"net/http"
	"rime-api/internal/auth"
	"rime-api/internal/mailer"
	"rime-api/internal/store"
	"rime-api/internal/store/cache"
	"time"

	"go.uber.org/zap"
)

type application struct {
	config        config
	i18n          *i18nConfig
	store         store.Storage
	cacheStore    cache.Storage
	logger        *zap.SugaredLogger
	mailer        mailer.Client
	authenticator auth.Authenticator
	router        http.Handler
}

func (app *application) serveHttp() error {
	app.routes()

	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      app.router,
		WriteTimeout: time.Second * 45,
		ReadTimeout:  time.Second * 20,
		IdleTimeout:  time.Minute,
	}

	app.logger.Infof("server has started at %s", app.config.addr)

	return server.ListenAndServe()
}
