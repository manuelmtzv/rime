package main

import (
	"rime-api/internal/auth"
	"rime-api/internal/db"
	"rime-api/internal/env"
	"rime-api/internal/mailer"
	"rime-api/internal/store"
	"rime-api/internal/store/cache"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func main() {
	bundle, localizer := initI18n()

	i18n := &i18nConfig{
		bundle:    bundle,
		localizer: localizer,
	}

	logger := zap.Must(zap.NewProduction()).Sugar()

	err := env.Load()
	if err != nil {
		logger.Panic(err)
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:password@localhost:5432/rime-db?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 25),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		redisdbCfg: redisConfig{
			addr:    env.GetString("REDIS_ADDR", "localhost:6379"),
			pw:      env.GetString("REDIS_PW", ""),
			db:      env.GetInt("REDIS_DB", 0),
			enabled: env.GetBool("REDIS_ENABLED", false),
		},
		mail: mailConfig{
			config: brevoConfig{
				apiKey:     env.GetString("BREVO_API_KEY", ""),
				partnerKey: env.GetString("BREVO_PARTNER_KEY", ""),
			},
		},
		auth: authConfig{
			jwt: jwtConfig{
				secret:         env.GetString("JWT_SECRET", "secret"),
				expires:        env.GetDuration("JWT_EXPIRES", time.Hour),
				refreshSecret:  env.GetString("JWT_REFRESH_SECRET", ""),
				refreshExpires: env.GetDuration("JWT_REFRESH_EXPIRES", time.Hour*24*7),
				issuer:         env.GetString("JWT_ISSUER", "rime-api"),
			},
		},
	}

	authenticator := auth.NewJWTAuthenticator(cfg.auth.jwt.secret, cfg.auth.jwt.refreshSecret, cfg.auth.jwt.issuer, cfg.auth.jwt.issuer)

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
	var rdb *redis.Client
	if cfg.redisdbCfg.enabled {
		rdb = cache.NewRedisClient(cfg.redisdbCfg.addr, cfg.redisdbCfg.pw, cfg.redisdbCfg.db)

		defer rdb.Close()
		logger.Infow("Redis connection established")
	}

	cacheStorage := cache.NewRedisStorage(rdb)

	app := &application{
		config:        cfg,
		i18n:          i18n,
		store:         store,
		cacheStore:    cacheStorage,
		logger:        logger,
		mailer:        mailer,
		authenticator: authenticator,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
