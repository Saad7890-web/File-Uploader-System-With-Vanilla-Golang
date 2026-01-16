package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/Saad7890-web/file-upload/internal/bootstrap"
	"github.com/Saad7890-web/file-upload/internal/infrastructure/config"
	"github.com/Saad7890-web/file-upload/internal/infrastructure/db"
)

func main() {
	
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	cfg := config.Load()

	database, err := db.NewPostgres(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	txManager := bootstrap.NewTransactionManager(database)

	_ = txManager

	log.Println("Application started successfully")
}
