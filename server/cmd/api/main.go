package main

import (
	"rime-api/internal/db"
	"rime-api/internal/env"
	"rime-api/internal/mailer"
	"rime-api/internal/store"

	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction()).Sugar()

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:password@localhost:5432/rime-db?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 25),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		mail: mailConfig{
			config: brevoConfig{
				apiKey:     env.GetString("BREVO_API_KEY", ""),
				partnerKey: env.GetString("BREVO_PARTNER_KEY", ""),
			},
		},
	}

	mailer := mailer.NewBrevo(cfg.mail.config.apiKey, cfg.mail.config.partnerKey)

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Panic(err)
	}

	defer db.Close()
	logger.Infow("Database connection established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
		mailer: mailer,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
