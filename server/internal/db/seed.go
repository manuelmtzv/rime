package db

import (
	"context"
	"database/sql"
	"log"
	"rime-api/internal/models"
	"rime-api/internal/store"
)

var usernames = [][]string{
	{"manuel.mtzv", "Manuel", "Martínez"},
	{"jose.garcia", "José", "García"},
	{"luis.lopez", "Luis", "López"},
	{"maria.perez", "María", "Pérez"},
	{"carmen.gomez", "Carmen", "Gómez"},
	{"ana.fernandez", "Ana", "Fernández"},
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(6)
	// tx, _ := db.BeginTx(ctx, nil)

	for _, u := range users {
		if err := store.Users.Create(ctx, u); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	log.Println("Seed completed")
}

func generateUsers(num int) []*models.User {
	users := make([]*models.User, num)
	for i := 0; i < num; i++ {
		u := &models.User{
			Username: usernames[i][0],
			Name:     usernames[i][1],
			Lastname: usernames[i][2],
			Password: "123456",
			Email:    usernames[i][0] + "@example.com",
		}
		users[i] = u
	}
	return users
}
