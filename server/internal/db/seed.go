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
	"<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p><p>Vivamus luctus urna sed urna ultricies ac tempor dui sagittis.</p><p>In condimentum facilisis porta. Sed nec diam eu diam mattis viverra.</p><br/><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p><p>Vivamus luctus urna sed urna ultricies ac tempor dui sagittis.</p><p>In condimentum facilisis porta. Sed nec diam eu diam mattis viverra.</p><br/><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p><p>Vivamus luctus urna sed urna ultricies ac tempor dui sagittis.</p><p>In condimentum facilisis porta. Sed nec diam eu diam mattis viverra.</p>",
	"<p>Cras sit amet libero eros. Fusce in fermentum velit.</p><p>Aliquam eget nisi scelerisque, placerat odio nec, condimentum erat.</p><p>Morbi eget posuere nisl. Vestibulum ante ipsum primis in faucibus orci luctus.</p>",
	"<p>Etiam luctus dapibus tortor et consequat. Nullam id dui sapien.</p><p>Pellentesque habitant morbi tristique senectus et netus et malesuada.</p><p>Nunc non turpis in turpis feugiat euismod nec vel metus.</p>",
	"<p>Integer placerat ipsum vel nibh tincidunt rhoncus.</p><p>Phasellus vel tincidunt lorem. Vestibulum at risus dolor.</p><p>Mauris sagittis tortor quis ultricies fermentum. Sed ac ante eget.</p>",
	"<p>Proin pharetra, risus ut lobortis elementum, massa erat commodo.</p><p>Nam eget placerat lorem, sit amet egestas mauris.</p><p>Curabitur tempor diam lacus, nec aliquet nisi rhoncus eget.</p>",
	"<p>Vivamus quis ultricies eros. Integer aliquet viverra nulla.</p><p>Aliquam erat volutpat. Pellentesque sed odio nibh.</p><p>Nulla malesuada lobortis nibh, ut tempus enim scelerisque id.</p>",
	"<p>Donec consectetur felis at massa venenatis, et fringilla dui gravida.</p><p>Maecenas posuere odio nec fermentum fringilla.</p><p>Ut dignissim ornare felis. Curabitur et consequat sapien.</p>",
}

var tagNames = []string{
	"Health",
	"Science",
	"Technology",
	"Art",
	"Poetry",
	"Philosophy",
	"Politics",
	"History",
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

	tags := generateTags(6)

	for _, t := range tags {
		if err := store.Tags.Create(ctx, t); err != nil {
			log.Println("Error creating tag:", err)
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

func generateTags(num int) []*models.Tag {
	tags := make([]*models.Tag, num)

	for i := 0; i < num; i++ {
		t := &models.Tag{
			Name: tagNames[i],
		}
		tags[i] = t
	}
	return tags
}

func randomNumber(min, max int) int {
	return min + rand.Intn(max-min)
}
