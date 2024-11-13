package main

import (
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type config struct {
	addr       string
	db         dbConfig
	mail       mailConfig
	auth       authConfig
	redisdbCfg redisConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type authConfig struct {
	jwt jwtConfig
}

type jwtConfig struct {
	secret         string
	expires        time.Duration
	refreshSecret  string
	refreshExpires time.Duration
	issuer         string
}

type mailConfig struct {
	config brevoConfig
}

type brevoConfig struct {
	apiKey     string
	partnerKey string
}

type redisConfig struct {
	addr    string
	pw      string
	db      int
	enabled bool
}

type i18nConfig struct {
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}
