package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	connection "github.com/xenyy/train-booking-api/db"
	"github.com/xenyy/train-booking-api/models"
	"github.com/xenyy/train-booking-api/server"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	db, err := connection.New(&connection.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		log.Fatal("failed to connect to postgresql database", err)
	}

	err = connection.SetupDatabase(db,
		&models.Ticket{},
	)

	if err != nil {
		log.Fatal("failed to setup tables", err)
	}

	r := server.StartServer(db)

	r.Run(":8000")
}
