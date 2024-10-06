package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	dbname   = os.Getenv("DB_NAME")
	password = os.Getenv("DB_PASSWORD")
	sslmode  = os.Getenv("DB_SSLMODE")
)

func New() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	fmt.Println("Connected to database")

	return db
}
