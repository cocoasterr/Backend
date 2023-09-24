package PG

import (
	"log"
	"os"

	model "github.com/cocoasterr/Backend/go_app/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PGConfigGorm() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to access .env!")
		return nil, err
	}

	db_udrl := os.Getenv("PG_DB_URL")
	db, err := gorm.Open(postgres.Open(db_udrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect Postgres!")
		return nil, err
	}
	db.AutoMigrate(&model.Product{})
	log.Println("Connected to Postgres!")
	return db, nil
}
