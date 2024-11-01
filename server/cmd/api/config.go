package main

import "time"

type config struct {
	addr string
	db   dbConfig
	mail mailConfig
	auth authConfig
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
	secret  string
	expires time.Duration
	issuer  string
}

type mailConfig struct {
	config brevoConfig
}

type brevoConfig struct {
	apiKey     string
	partnerKey string
}
