package main

import (
	"log"
	"rime-api/internal/db"
	"rime-api/internal/env"
	"rime-api/internal/store"
)

func main() {
	add := env.GetString("DB_ADDR", "postgres://postgres:password@localhost:5432/rime-db?sslmode=disable")
	conn, err := db.New(add, 25, 25, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store, conn)
}
