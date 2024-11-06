package db

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
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

var writtingContents = []string{
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	"Vivamus auctor, nunc nec lacinia ultricies, nunc nunc.",
	"Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae.",
	"Nullam nec nunc nec nunc.",
	"Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae.",
	"Nullam nec nunc nec nunc.",
	"Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae.",
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

	writings := generateWritings(6, users)

	for _, w := range writings {
		if err := store.Writings.Create(ctx, w); err != nil {
			log.Println("Error creating writting:", err)
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

func generateWritings(num int, users []*models.User) []*models.Writing {

	writings := make([]*models.Writing, num)

	for i := 0; i < num; i++ {
		w := &models.Writing{
			Title:    "Title",
			Type:     "poetry",
			Content:  writtingContents[i],
			AuthorID: users[randomNumber(0, len(users))].ID,
		}
		writings[i] = w
	}
	return writings
}

func randomNumber(min, max int) int {
	return min + rand.Intn(max-min)
}
